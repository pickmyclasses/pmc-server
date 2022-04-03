package dao

import (
	"errors"
	"gorm.io/gorm"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetUserCourseHistoryList(userID int64) ([]model.UserCourseHistory, error) {
	var courseHistoryList []model.UserCourseHistory
	res := postgres.DB.Where("user_id = ?", userID).Find(&courseHistoryList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return courseHistoryList, nil
}

func GetUserCourseHistoryByID(userID, courseID int64) (*model.UserCourseHistory, error) {
	var courseHistory model.UserCourseHistory
	res := postgres.DB.Where("user_id = ? and course_id = ?", userID, courseID).First(&courseHistory)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &courseHistory, nil
}

func CreateSingleUserCourseHistory(userID, courseID int64, semesterID int32, professorName string) error {
	history := model.UserCourseHistory{
		UserID:     userID,
		CourseID:   courseID,
		SemesterID: semesterID,
		ProfessorName:    professorName,
	}
	res := postgres.DB.Create(&history)
	if res.Error != nil {
		return shared.InternalErr{}
	}
	return nil
}

func DeleteSingleUserCourseHistory(userID, courseID int64, semesterID int32) error {
	var history model.UserCourseHistory
	res := postgres.DB.Where("user_id = ? and course_id = ? and semester_id = ?", userID, courseID, semesterID).First(&history)
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

func CheckIfCourseInUserCourseHistory(userID, courseID int64) (bool, error) {
	var history model.UserCourseHistory
	res := postgres.DB.Where("user_id = ? and course_id = ?", userID, courseID).First(&history)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, shared.InternalErr{}
	}
	return true, nil
}
