package schedule

import (
	"errors"

	"pmc_server/init/postgres"
	"pmc_server/model"
)

func CheckIfUserExist(userID int64) (bool, error) {
	var user model.User
	res := postgres.DB.Where("id = ?", userID).Find(&user)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func CheckIfClassExist(classID int64) (bool, error) {
	var class model.Class
	res := postgres.DB.Where("id = ?", classID).Find(&class)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return true, nil
	}
	return true, nil
}

func CheckIfScheduleExist(classID, userID, semesterID int64) (bool, error) {
	var schedule model.Schedule
	res := postgres.DB.
		Where("user_id = ? and class_id = ? and semester_id = ?", userID, classID, semesterID).
		Find(&schedule)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func CreateSchedule(classID, userID, semesterID int64) error {
	schedule := &model.Schedule{
		UserID:     userID,
		ClassID:    classID,
		SemesterID: semesterID,
	}
	res := postgres.DB.Create(&schedule)
	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("create schedule failed")
	}
	return nil
}

func GetScheduleByUserID(userID int64) ([]model.Schedule, error) {
	var schedule []model.Schedule
	res := postgres.DB.Where("user_id = ?", userID).Find(&schedule)
	if res.Error != nil {
		return nil, res.Error
	}
	return schedule, nil
}

func DeleteUserSchedule(userID, semesterID, classID int64) error {
	var schedule model.Schedule
	res := postgres.DB.
		Where("user_id = ? AND class_id = ? AND semester_id = ?", userID, classID, semesterID).
		First(&schedule)

	if res.RowsAffected == 0 {
		return errors.New("no user schedule found")
	}

	if res.Error != nil {
		return errors.New("internal error occurred while fetching schedule")
	}

	res = postgres.DB.
		Where("user_id = ? AND class_id = ? AND semester_id = ?", userID, classID, semesterID).
		Delete(&schedule)

	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("delete user seclude failed")
	}
	return nil
}
