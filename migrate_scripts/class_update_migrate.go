package migrate

import (
	"fmt"
	"pmc_server/model"
	"pmc_server/shared"
)

func UpdateClasses() {
	err, db := InitDB()
	if err != nil {
		panic(err)
	}

	var classList []model.Class
	res := db.Find(&classList)
	if res.Error != nil {
		panic(res.Error.Error())
	}

	for _, class := range classList {
		if class.StartTime == "" && class.EndTime == "" && class.Component == "In Person" {
			if res := db.Model(&class).Update("component", "Online"); res.Error != nil {
				panic(err)
			}
		}
		if class.StartTime != "" && class.EndTime != "" && class.Component == "Online" {
			if res := db.Model(&class).Update("component", "In Person"); res.Error != nil {
				panic(err)
			}
		}
	}
}

func UpdateOfferTime() {
	err, db := InitDB()
	if err != nil {
		panic(err)
	}

	var classList []model.Class
	res := db.Find(&classList)
	if res.Error != nil {
		panic(res.Error.Error())
	}

	oddList := make([]int64, 0)

	for _, class := range classList {
		if class.StartTime != "" {
			newStartTime, err := shared.ConvertTimestamp(class.StartTime)
			if err != nil {
				fmt.Printf("failed to update class %d\n", class.ID)
			}
			if res := db.Model(&class).Update("start_time_float", newStartTime); res.Error != nil {
				panic(res.Error.Error())
			}
		}

		if class.EndTime != "" {
			newEndTime, err := shared.ConvertTimestamp(class.EndTime)
			if err != nil {
				oddList = append(oddList, class.ID)
			}
			if res := db.Model(&class).Update("end_time_float", newEndTime); res.Error != nil {
				panic(res.Error.Error())
			}
		}
	}
	fmt.Println(oddList)
}
