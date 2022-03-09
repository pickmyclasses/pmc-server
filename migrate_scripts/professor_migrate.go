package migrate

import (
	"fmt"
	"pmc_server/model"
	"strings"

	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() (err error, db *gorm.DB) {
	dsn := "host=pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com user=admin1 password=admin123 dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return
}

func Professors() {
	err, db := InitDB()
	if err != nil {
		fmt.Errorf("error when init database %+v", err)
	}

	var classes []model.Class
	res := db.Find(&classes)
	if res.RowsAffected == 0 || res.Error != nil {
		fmt.Errorf("error when fetching courses")
	}

	professors := make([]string, 0)
	mapping := make(map[string]bool)
	for _, course := range classes {
		if course.Instructors != "" {
			if _, exist := mapping[course.Instructors]; !exist {
				mapping[course.Instructors] = true
				professors = append(professors, course.Instructors)
			}
		}
	}

	for _, prof := range professors {
		professor := model.Professor{Name: prof, CollegeID: 1}
		res := db.Create(&professor)
		if res.Error != nil || res.RowsAffected == 0 {
			panic("error when creating professor")
		}
	}
}

func ProfessorIDs() {
	err, db := InitDB()
	if err != nil {
		fmt.Errorf("error when init database %+v", err)
	}

	var classes []model.Class
	res := db.Find(&classes)
	if res.RowsAffected == 0 || res.Error != nil {
		fmt.Errorf("error when fetching courses")
	}

	var professors []model.Professor
	res = db.Find(&professors)
	if res.RowsAffected == 0 || res.Error != nil {
		fmt.Errorf("error when fetching professors")
	}

	for _, class := range classes {
		for _, professor := range professors {
			if strings.TrimSpace(class.Instructors) == strings.TrimSpace(professor.Name) {
				class.InstructorID = int32(professor.ID)
				res = db.Save(&class)
				if res.RowsAffected == 0 || res.Error != nil {
					fmt.Errorf("error when updating ID field")
				}
			}
		}
	}

}
