package migrate

import (
	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math/rand"
	"pmc_server/model"
	"time"
)

func PopulateRandomData() error {
	dsn := "host=pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com user=admin1 password=pickmyclassesdotcom dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	var course model.Course
	res := db.Where("id =? ", 31826).Find(&course)
	if res.Error != nil {
		return res.Error
	}

	var semesterList []model.Semester
	res = db.Find(&semesterList)
	if res.Error != nil {
		return res.Error
	}

	for _, semester := range semesterList {
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(140-20) + 20
		coursePop := model.CoursePopularity{
			CourseID:   course.ID,
			Popularity: int32(randomNum),
			SemesterID: int32(semester.ID),
		}
		res = db.Create(&coursePop)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
