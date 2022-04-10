package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetSemesterByID(semesterID int32) (*model.Semester, error) {
	var semester model.Semester
	res := postgres.DB.Where("id = ?", semesterID).First(&semester)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &semester, nil
}

func GetSemesterByName(season string, year int32) (*model.Semester, error) {
	var semester model.Semester
	res := postgres.DB.Where("season = ? and year = ?", season, year)
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
