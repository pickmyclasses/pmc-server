package dao

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetUserCourseHistory(userID int64)  ([]model.UserCourseHistory, error) {
	var courseHistoryList []model.UserCourseHistory
	res := postgres.DB.Where("user_id = ?", userID).Find(&courseHistoryList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return courseHistoryList, nil
}

func CreateSingleUserCourseHistory(userID, courseID, classID int64, semesterID int32) error{
	history := model.UserCourseHistory{
		UserID:     userID,
		CourseID:   courseID,
		SemesterID: semesterID,
		ClassID: classID,
	}
	res := postgres.DB.Create(&history)
	if res.Error != nil {
		return shared.InternalErr{}
	}
	return nil
}

func DeleteSingleUserCourseHistory(userID, courseID int64, semesterID int32) error {
	var history model.UserCourseHistory
	res := postgres.DB.Where("user_id = ?, course_id = ?, semester_id = ?", userID, courseID, semesterID).First(&history)
	if res.Error != nil {
		return shared.InternalErr{}
	}
	if res.RowsAffected == 0 {
		return shared.ContentNotFoundErr{}
	}
	res = postgres.DB.Delete(&history)
	if res.Error != nil {
		return shared.InternalErr{}
	}
	return nil
}
