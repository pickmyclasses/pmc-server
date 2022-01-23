package main

import (
	"errors"
	"fmt"
	model "pmc_server/model"
	"pmc_server/utils"
	"strings"

	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init() (err error, db *gorm.DB) {
	dsn := "host=localhost user=postgres password=admin123 dbname=pmc port=5432 sslmode=disable"
	db, err = gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return
}

func UpdateCourseAndClasses() error {
	err, db := Init()
	if err != nil {
		return err
	}

	var courses []model.Course
	res := db.Find(&courses)
	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("error when fetching the courses")
	}

	var classes []model.Class
	classRes := db.Find(&classes)
	if classRes.RowsAffected == 0 || classRes.Error != nil {
		return errors.New("error when fetching the classes")
	}

	var classDeletion model.Course
	delRes := db.Where("catalog_course_name = ?", "").Delete(&classDeletion)
	if delRes.Error != nil {
		return errors.New("error when deleting course data")
	}

	for _, class := range classes {
		letterClass, numberClass := utils.GetLetterInfo(class.CourseCatalogName)
		for _, course := range courses {
			letter, number := utils.ParseString(course.CatalogCourseName, true)
			if letter == strings.ToUpper(letterClass) && numberClass == number {
				res := db.Model(&class).Update("course_id", course.ID)
				if res.RowsAffected == 0 || res.Error != nil {
					return errors.New("unable to update")
				}
			}
		}
	}

	return nil
}

func main() {
	err := UpdateCourseAndClasses()
	if err != nil {
		fmt.Println(err)
	}
}
