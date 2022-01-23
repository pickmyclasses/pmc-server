package main

import (
	"errors"
	"fmt"

	model "pmc_server/models"

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
		return errors.New("error when fetching teh classes")
	}

	noID := 0
	for _, course := range courses {
		if course.CatalogCourseName == "" {
			noID++
		}
	}
	fmt.Println(noID)

	return nil
}

func main() {
	err := UpdateCourseAndClasses()
	if err != nil {
		fmt.Println(err)
	}
}
