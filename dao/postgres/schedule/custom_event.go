package schedule

import (
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
	color string, days []int64, startTime, endTime int32) error {
	event := model.CustomEvent{
		Title:       title,
		Description: description,
		Color:       color,
		Days:        days,
		StartTime:   startTime,
		EndTime:     endTime,
		UserID:      userID,
		SemesterID:  semesterID,
	}
	res := postgres.DB.Create(&event)

	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}

func CheckIfCustomEventExist(id int64) (bool, error) {
	var event model.CustomEvent
	res := postgres.DB.Where("id = ?", id).Find(&event)
	if res.Error != nil {
		return false, shared.InternalErr{}
	}

	return res.RowsAffected == 0, nil
}

func UpdateCustomEventByID(userID, semesterID int64, title, description,
	color string, days []int64, startTime, endTime int32) error {
	event := &model.CustomEvent{
		Title:       title,
		Description: description,
		Color:       color,
		Days:        days,
		StartTime:   startTime,
		EndTime:     endTime,
		UserID:      userID,
		SemesterID:  semesterID,
	}
	res := postgres.DB.Save(&event)
	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}
