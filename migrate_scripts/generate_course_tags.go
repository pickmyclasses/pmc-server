package migrate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"pmc_server/model"

	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Response struct {
	Keyword    map[string]int     `json:"keyword"`
	Topic      map[string]float32 `json:"topic"`
	Version    string             `json:"version"`
	Author     string             `json:"author"`
	Email      string             `json:"email"`
	ResultCode string             `json:"result_code"`
	ResultMsg  string             `json:"result_msg"`
}

func GenerateCourseAssociatedTags() error {
	dsn := "host=pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com user=admin1 password=pickmyclassesdotcom dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return err
	}

	var courseList []model.Course
	dbRes := db.Find(&courseList)
	if dbRes.Error != nil {
		return dbRes.Error
	}

	host := "twinword-topic-tagging.p.rapidapi.com"
	key := "aASFE865ImTtakxSwxZYWXFO8Jmm8ZtG955HM3hNiNIo7bU6IvLtFgvhY2f8jgSamzKlXfHnGISNKiv9eTo/pQ=="

	for _, course := range courseList {
		text := course.Description
		text = strings.ReplaceAll(strings.TrimSpace(text), " ", "%20")
		if text == "" {
			continue
		}
		uri := fmt.Sprintf("https://api.twinword.com/api/topic/generate/latest/?text=%s", text)

		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			return err
		}
		req.Header.Add("X-RapidAPI-Host", host)
		req.Header.Add("X-RapidAPI-Key", key)

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()

		var result Response
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return err
		}

		if result.ResultCode != "200" {
			return fmt.Errorf("unable to fetch data %+v\n", result.ResultMsg)
		}
		for k, v := range result.Topic {
			associatedTag := model.AssociatedTag{CourseID: course.ID, Content: k, Weight: v}
			createRes := db.Create(&associatedTag)
			if createRes.Error != nil {
				return createRes.Error
			}
		}

		for k, v := range result.Keyword {
			associatedTag := model.AssociatedTag{CourseID: course.ID, Content: k, Weight: float32(v)}
			createRes := db.Create(&associatedTag)
			if createRes.Error != nil {
				return createRes.Error
			}
		}
	}

	return nil
}
