package course

import (
	"errors"
	"pmc_server/init/postgres"

	model "pmc_server/models"
	"pmc_server/utils"
)

func GetCourses(pn, pSize int) ([]model.Course, int64) {
	var courses []model.Course
	total := postgres.DB.Find(&courses).RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Find(&courses)

	return courses, total
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

func GetClassListByCourseID(id, pn, pSize int) (*[]model.Class, int64) {
	var classes []model.Class
	total := postgres.DB.Where("course_id = ?", id).Find(&classes).RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Find(&classes)
	return &classes, total
}
