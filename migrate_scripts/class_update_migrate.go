package migrate

import (
	"pmc_server/model"
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

