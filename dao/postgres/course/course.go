package course

import (
	"errors"
	"pmc_server/shared"

	"pmc_server/init/postgres"
	"pmc_server/model"
)

func GetCourses(pn, pSize int) ([]model.Course, error) {
	var courses []model.Course
	res := postgres.DB.Scopes(shared.Paginate(pn, pSize)).Find(&courses)
	if res.Error != nil {
		return nil, res.Error
	}
	return courses, nil
}

func GetCourseTotal() (int64, error) {
	var total int64
	res := postgres.DB.Model(&model.Course{}).Count(&total)
	if res.Error != nil {
		return -1, res.Error
	}

	return total, nil
}

func GetCourseByID(id int) (*model.Course, error) {
	var course model.Course
	result := postgres.DB.Where("id = ?", id).First(&course)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no course info found")
	}
	return &course, nil
}

func GetClassListByCourseID(id int) (*[]model.Class, int64) {
	var classes []model.Class
	res := postgres.DB.Where("course_id = ?", id).Find(&classes)
	return &classes, res.RowsAffected
}
