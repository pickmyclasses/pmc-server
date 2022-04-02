package dao

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetUserCourseHistory(userID int64)  ([]model.UserCourseHistory, error) {
	return nil, nil
}

func CreateSingleUserCourseHistory(userID, courseID int64, semesterID int32) error{
	history := model.UserCourseHistory{
		UserID:     userID,
		CourseID:   courseID,
		SemesterID: semesterID,
	}
	res := postgres.DB.Create(&history)
	if res.Error != nil {
		return shared.InternalErr{}
	}
	return nil
}

func DeleteSingleUserCourseHistory(userID, courseID int64, semesterID int32) error {
	return nil
}
