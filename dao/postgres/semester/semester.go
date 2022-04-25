package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
	"strconv"
)

func GetSemesterByID(semesterID int32) (*model.Semester, error) {
	var semester model.Semester
	res := postgres.DB.Where("id = ?", semesterID).First(&semester)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &model.Semester{}, nil
		}
		return nil, shared.InternalErr{}
	}
	return &semester, nil
}

func GetSemesterByName(season string, year int32) (*model.Semester, error) {
	yearStr := strconv.Itoa(int(year))
	var semester model.Semester
	res := postgres.DB.Where("season = ? and year = ?", season, yearStr).First(&semester)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to fetch the semester %s %d", season, year),
		}
	}
	return &semester, nil
}

func GetSemesterListByCollegeID(collegeID int32) ([]model.Semester, error) {
	var semesterList []model.Semester
	res := postgres.DB.Where("college_id = ?", collegeID).Find(&semesterList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []model.Semester{}, nil
		}
	}
	return semesterList, nil
}
