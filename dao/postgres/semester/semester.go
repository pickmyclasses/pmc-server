package dao

import (
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
