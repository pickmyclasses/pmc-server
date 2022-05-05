package dao

import (
	"errors"
	"gorm.io/gorm"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

// GetCollegeList gives the entire list of the colleges in the database
func GetCollegeList() ([]model.College, error) {
	var collegeList []model.College
	res := postgres.DB.Find(&collegeList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return collegeList, nil
}

// GetCollegeSemesterList gives the semester list in a college
func GetCollegeSemesterList(collegeID int32) ([]model.Semester, error) {
	var semesterList []model.Semester
	res := postgres.DB.Where("college_id = ?", collegeID).Find(&semesterList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return semesterList, nil
}

// GetCollegeByID gives a college entity with the given college ID
func GetCollegeByID(collegeID int32) (*model.College, error) {
	var college model.College
	res := postgres.DB.Where("id = ?", collegeID).First(&college)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &model.College{}, nil
		}
		return nil, shared.InternalErr{}
	}
	return &college, nil
}
