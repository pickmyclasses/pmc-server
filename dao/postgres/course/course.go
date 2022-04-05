package dao

import (
	"pmc_server/shared"
	"strings"

	"pmc_server/init/postgres"
	"pmc_server/model"
)

func GetCourses(pn, pSize int) ([]model.Course, error) {
	var courses []model.Course
	res := postgres.DB.Scopes(shared.Paginate(pn, pSize)).Find(&courses)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return courses, nil
}

func GetCourseTotal() (int64, error) {
	var total int64
	res := postgres.DB.Model(&model.Course{}).Count(&total)
	if res.Error != nil {
		return -1, shared.InternalErr{}
	}

	return total, nil
}

func GetCourseByID(id int) (*model.Course, error) {
	var course model.Course
	result := postgres.DB.Where("id = ?", id).First(&course)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	if result.RowsAffected == 0 {
		return nil, shared.ContentNotFoundErr{}
	}
	return &course, nil
}

func GetClassListByCourseID(id int) (*[]model.Class, int64) {
	var classes []model.Class
	res := postgres.DB.Where("course_id = ?", id).Find(&classes)
	return &classes, res.RowsAffected
}

func GetCourseByCatalogName(catalogName string) (*model.Course, error) {
	var course model.Course
	result := postgres.DB.Where("catalog_course_name", strings.TrimSpace(catalogName)).First(&course)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &course, nil
}
