package dao

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetCollegeList() ([]model.College, error) {
	var collegeList []model.College
	res := postgres.DB.Find(&collegeList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return collegeList, nil
}
func GetCollegeSemesterList(collegeID int32) ([]model.Semester, error) {
	var semesterList []model.Semester
	res := postgres.DB.Where("college_id = ?", collegeID).Find(&semesterList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return semesterList, nil
}

func GetCollegeByID(collegeID int32) (*model.College, error) {
	var college model.College
	res := postgres.DB.Where("id = ?", collegeID).First(&college)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &college, nil
}
