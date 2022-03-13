package dao

import (
	"fmt"
	"sort"

	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetClasses(pn, pSize int) (*[]model.Class, int64) {
	var classes []model.Class
	total := postgres.DB.Find(&classes).RowsAffected
	postgres.DB.Scopes(shared.Paginate(pn, pSize)).Find(&classes)

	return &classes, total
}

func GetClassInfoByID(id int) (*model.Class, error) {
	var class model.Class
	result := postgres.DB.Where("id = ?", id).First(&class)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &class, nil
}

func GetClassByCourseID(courseID int64) (*[]model.Class, error) {
	var classes []model.Class
	result := postgres.DB.Where("course_id = ?", courseID).Find(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

func GetClassListByComponent(components []string) (*[]model.Class, error) {
	var classes []model.Class
	sql := "select * from class where component = "
	for i, c := range components {
		if i == len(components)-1 {
			sql += fmt.Sprintf("'%s'", c)
		} else {
			sql += fmt.Sprintf("'%s' or component = ", c)
		}
	}

	result := postgres.DB.Raw(sql).Scan(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

func GetClassListByOfferDate(offerDates []int) (*[]model.Class, error) {
	// sort the dates first to convert to the correct format
	sort.Slice(offerDates, func(i, j int) bool {
		return offerDates[i] < offerDates[j]
	})

	dates := shared.ConvertSliceToDateString(offerDates)

	var classes []model.Class
	result := postgres.DB.Where("offer_date = ?", dates).Find(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

//func GetClassListByTimeslot(startTime, endTime float32) (*[]model.Class, error) {
//
//}

//func GetClassListByProfessorNames(professorNames []string) (*[]model.Class, error) {
//
//}
