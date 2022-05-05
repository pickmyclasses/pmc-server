package migrate

import (
	"context"
	"fmt"
	"log"
	"os"
	"pmc_server/shared"
	"strconv"
	"strings"

	"pmc_server/model"
	esModel "pmc_server/model/es"

	"github.com/olivere/elastic/v7"
	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CourseEs() {
	url := "http://34.227.17.136:9200"
	logger := log.New(os.Stdout, "pmc", log.LstdFlags)
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		//elastic.SetBasicAuth(username, password),
		elastic.SetTraceLog(logger),
	)

	if err != nil {
		fmt.Println("failed to init es")
		return
	}

	dsn := "host=pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com user=admin1 password=pickmyclassesdotcom dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("failed to init postgres")
		return
	}

	var courses []model.Course
	res := db.Find(&courses)
	if res.Error != nil || res.RowsAffected == 0 {
		fmt.Println("failed to fetch courses")
		return
	}

	for _, course := range courses {
		maxCredit, err := strconv.ParseFloat(course.MaxCredit, 32)
		if err != nil {
			fmt.Println("failed to convert max credit")
			return
		}

		minCredit, err := strconv.ParseFloat(course.MinCredit, 32)
		if err != nil {
			fmt.Println("failed to convert min credit")
			return
		}

		var rating model.OverAllRating
		_ = db.Where("course_id = ?", course.ID).First(&rating)

		letter, number := shared.ParseString(course.CatalogCourseName, false)

		var classes []model.Class
		_ = db.Where("course_id = ?", course.ID).Find(&classes)

		var tags []model.Tag
		_ = db.Where("course_id = ?", course.ID).Find(&tags)
		tagNameList := make([]string, 0)
		for _, t := range tags {
			tagNameList = append(tagNameList, strings.TrimSpace(t.Name))
		}

		classNestedList := make([]esModel.ClassNested, 0)

		for _, cls := range classes {
			classNested := esModel.ClassNested{
				Semester:    cls.Semester,
				Year:        2022,
				WaitList:    cls.WaitList,
				OfferDate:   cls.OfferDate,
				StartTime:   cls.StartTimeFloat,
				EndTime:     cls.EndTimeFloat,
				Instructors: cls.Instructors,
			}
			classNestedList = append(classNestedList, classNested)
		}

		esCourse := esModel.Course{
			ID:                 course.ID,
			DesignationCatalog: course.DesignationCatalog,
			Title:              course.Title,
			Description:        course.Description,
			CatalogCourseName:  fmt.Sprintf("%s %s", letter, number),
			CatalogName:        letter,
			Prerequisites:      course.Prerequisites,
			Component:          course.Component,
			MaxCredit:          float32(maxCredit),
			MinCredit:          float32(minCredit),
			IsHonor:            course.IsHonor,
			FixedCredit:        course.FixedCredit,
			Rating:             rating.OverAllRating,
			HasOffering:        len(classes) != 0,
			Classes:            classNestedList,
			Tags:               tagNameList,
		}

		_, err = client.Index().Index(esCourse.GetIndexName()).BodyJson(esCourse).Id(strconv.Itoa(int(course.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
