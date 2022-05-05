package dao

import (
	"github.com/lib/pq"

	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetCustomEventByUserID(userID int64) ([]model.CustomEvent, error) {
	var events []model.CustomEvent
	res := postgres.DB.Where("user_id = ?", userID).Find(&events)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return events, nil
}

func CreateCustomEventByUserID(userID, semesterID int64, title, description,
	color string, days []int64, startTime, endTime int32, kind string) error {
	event := model.CustomEvent{
		Title:       title,
		Description: description,
		Color:       color,
		Days:        days,
		StartTime:   startTime,
		EndTime:     endTime,
		UserID:      userID,
		SemesterID:  semesterID,
		Kind:        kind,
	}
	res := postgres.DB.Create(&event)

	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}

func CheckIfCustomEventExist(id int64) (bool, error) {
	var event model.CustomEvent
	res := postgres.DB.Where("id = ?", id).First(&event)
	if res.Error != nil {
		return false, shared.InternalErr{}
	}

	return res.RowsAffected != 0, nil
}

func UpdateCustomEventByID(eventID, userID, semesterID int64, title, description,
	color string, days []int64, startTime, endTime int32, kind string) error {

	res := postgres.DB.Model(&model.CustomEvent{}).Where("id = ?", eventID).Updates(map[string]interface{}{
		"title":       title,
		"description": description,
		"color":       color,
		"days":        pq.Int64Array(days),
		"start_time":  startTime,
		"end_time":    endTime,
		"kind":        kind,
	})
	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}

func DeleteCustomEvent(id int64) error {
	var event model.CustomEvent
	res := postgres.DB.
		Where("id = ?", id).
		First(&event)

	if res.RowsAffected == 0 {
		return shared.ContentNotFoundErr{}
	}

	if res.Error != nil {
		return shared.InternalErr{}
	}

	res = postgres.DB.
		Delete(&event)

	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}
