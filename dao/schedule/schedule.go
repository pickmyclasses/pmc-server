package schedule

import (
	"errors"
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
