package schedule

import (
	"errors"

	"go.uber.org/zap"

	"pmc_server/init/postgres"
	"pmc_server/model"
)

func PostUserSchedule(userID, classID, semesterID int64) error {
	var user model.User
	res := postgres.DB.Where("id = ?", userID).First(&user)
	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("user info not found")
	}
	var course model.Course
	res = postgres.DB.Where("id = ?", classID).First(&course)
	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("course info not found")
	}
	// TODO: fix the semester thing, check if it exist

	schedule := &model.Schedule{
		UserID:     userID,
		ClassID:    classID,
		SemesterID: 1,
	}
	res = postgres.DB.Create(&schedule)
	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("create schedule failed")
	}
	return nil
}

func GetUserSchedule(userID int64) (*model.Schedule, error) {
	var schedule model.Schedule
	res := postgres.DB.Where("user_id = ?", userID).First(&schedule)
	if res.RowsAffected == 0 || res.Error != nil {
		return nil, errors.New("no user schedule found")
	}
	return &schedule, nil
}

func DeleteUserSchedule(userID, semesterID, classID int64) error {
	var schedule model.Schedule
	res := postgres.DB.Where("user_id = ? AND class_id = ? AND semester_id = ?", userID, classID, semesterID).First(&schedule)
	if res.RowsAffected == 0 {
		return errors.New("no user schedule found")
	}

	if res.Error != nil {
		zap.L().Error("Error when fetching schedule")
		return errors.New("internal error occurred while fetching schedule")
	}

	res = postgres.DB.Where("user_id = ? AND class_id = ? AND semester_id = ?", userID, classID, semesterID).Delete(&schedule)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("delete user seclude failed")
	}
	return nil
}
