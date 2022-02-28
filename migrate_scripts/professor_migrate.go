package migrate

import (
	"fmt"
	"pmc_server/model"

	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() (err error, db *gorm.DB) {
	dsn := "host=localhost user=postgres password=admin123 dbname=pmc port=5432 sslmode=disable"
	db, err = gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return
}

func migrateProfessors() {
	err, db := InitDB()
	if err != nil {
		fmt.Errorf("error when init database %+v", err)
	}

	var courses []model.Course
	res := db.Find(&courses)
	if res.RowsAffected == 0 || res.Error != nil {
		fmt.Errorf("error when fetching courses")
	}

}
