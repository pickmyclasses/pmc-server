package migrate

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"pmc_server/model"
	"pmc_server/model/es"
	"pmc_server/shared"

	"github.com/olivere/elastic/v7"
	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Classes() {
	dsn := "host=pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com user=admin1 password=admin123 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	url := "https://search-pmc-search-jvq4ibvtwwkfg5kukvsexmll3q.us-east-1.es.amazonaws.com/"
	username := "admin1"
	password := "Admin123!"
	logger := log.New(os.Stdout, "pmc", log.LstdFlags)
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth(username, password),
		elastic.SetTraceLog(logger),
	)

	if err != nil {
		panic(err)
	}

	var classList []model.Class
	res := db.Find(&classList)
	if res.Error != nil {
		panic(res.Error.Error())
	}

	fmt.Printf("length of classList = %d\n", len(classList))

	for _, class := range classList {
		var isOnline bool
		var isInPerson bool
		var isIVC bool
		var isHybrid bool
		if strings.Contains(strings.ToLower(class.Component), "online") {
			isOnline = true
		}
		if strings.Contains(strings.ToLower(class.Component), "person") {
			isInPerson = true
		}
		if strings.Contains(strings.ToLower(class.Component), "ivc") {
			isIVC = true
		}
		if strings.Contains(strings.ToLower(class.Component), "hybrid") {
			isHybrid = true
		}

		offerDates := convertOfferDate(class.OfferDate)

		var translatedStartTime float32
		if class.StartTime != "" {
			translatedStartTime, _ = shared.ConvertTimestamp(class.StartTime)
		}

		var translatedEndTime float32
		if class.EndTime != "" {
			translatedEndTime, _ = shared.ConvertTimestamp(class.EndTime)
		}

		esClass := es.Class{
			ID:         class.ID,
			CourseID:   class.CourseID,
			IsOnline:   isOnline,
			IsInPerson: isInPerson,
			IsIVC:      isIVC,
			IsHybrid:   isHybrid,
			OfferDates: offerDates,
			StartTime:  translatedStartTime,
			EndTime:    translatedEndTime,
			Professors: class.Instructors,
		}

		_, err = client.Index().Index("class").BodyJson(esClass).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func convertOfferDate(offerDates string) []int {
	mapping := make(map[string]int)
	mapping["mo"] = 1
	mapping["tu"] = 2
	mapping["we"] = 3
	mapping["th"] = 4
	mapping["fr"] = 5

	offerLower := []rune(strings.ToLower(offerDates))
	res := make([]int, 0)
	curStr := ""
	for _, s := range offerLower {
		if s == '-' || s == ' ' {
			continue
		}
		curStr = curStr + string(s)
		if num, found := mapping[curStr]; found {
			res = append(res, num)
			curStr = ""
		}
	}
	return res
}
