package dao

import (
	"errors"
	"gorm.io/gorm"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func CheckIfUserExist(userID int64) (bool, error) {
	var user model.User
	res := postgres.DB.Where("id = ?", userID).Find(&user)
	if res.Error != nil {
		return false, shared.InternalErr{}
	}
	if res.RowsAffected == 0 {
		return false, shared.ContentNotFoundErr{}
	}
	return true, nil
}

func CheckIfClassExist(classID int64) (bool, error) {
	var class model.Class
	res := postgres.DB.Where("id = ?", classID).Find(&class)
	if res.Error != nil {
		return false, shared.InternalErr{}
	}
	if res.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func CheckIfScheduleExist(classID, userID, semesterID int64) (int64, error) {
	var schedule model.Schedule
	res := postgres.DB.
		Where("user_id = ? and class_id = ? and semester_id = ?", userID, classID, semesterID).
		Find(&schedule)
	if res.Error != nil {
		return -1, shared.InternalErr{}
	}
	if res.RowsAffected == 0 {
		return -1, nil
	}
	return schedule.ID, nil
}

func CreateSchedule(classID, userID, semesterID int64) error {
	schedule := &model.Schedule{
		UserID:     userID,
		ClassID:    classID,
		SemesterID: semesterID,
	}
	res := postgres.DB.Create(&schedule)
	if res.RowsAffected == 0 || res.Error != nil {
		return shared.InternalErr{}
	}
	return nil
}

func GetScheduleByUserID(userID int64) ([]model.Schedule, error) {
	var schedule []model.Schedule
	res := postgres.DB.Where("user_id = ?", userID).Find(&schedule)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []model.Schedule{}, nil
		}
		return nil, shared.InternalErr{}
	}
	return schedule, nil
}

func UpdateScheduleByID(id int64, classID, semesterID int64) error {
	res := postgres.DB.Model(&model.Schedule{}).Where("id = ?", id).
		Update("class_id", classID).
		Update("semester_id", semesterID)
	if res.Error != nil {
		return shared.InternalErr{}
	}
	return nil
}

func DeleteUserSchedule(userID, classID int64) error {
	var schedule model.Schedule
	// TODO: fix the semester ID
	res := postgres.DB.
		Where("user_id = ? and class_id = ? and semester_id = ?", userID, classID, 2).
		First(&schedule)

	if res.RowsAffected == 0 {
		return shared.ContentNotFoundErr{}
	}

	if res.Error != nil {
		return shared.InternalErr{}
	}

	res = postgres.DB.
		Delete(&schedule)

	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}
