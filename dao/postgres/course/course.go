package dao

import (
	"pmc_server/shared"
	"strings"

	"pmc_server/init/postgres"
	"pmc_server/model"
)

// GetCourseList gives the list of the courses
// this API should be used with page numbers and page size for pagination
func GetCourseList(pn, pSize int) ([]model.Course, error) {
	var courses []model.Course
	res := postgres.DB.Scopes(shared.Paginate(pn, pSize)).Find(&courses)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return courses, nil
}

// GetCourseTotal gives the total courses in database
func GetCourseTotal() (int64, error) {
	var total int64
	res := postgres.DB.Model(&model.Course{}).Count(&total)
	if res.Error != nil {
		return -1, shared.InternalErr{}
	}

	return total, nil
}

// GetCourseByID gets a course entity with the given ID
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

// GetClassListByCourseID gets the class list associated with the course with the given ID
func GetClassListByCourseID(id int) (*[]model.Class, int64) {
	var classes []model.Class
	res := postgres.DB.Where("course_id = ?", id).Find(&classes)
	return &classes, res.RowsAffected
}

// GetCourseByCatalogName gets a course entity with the given catalog name (eg, CS4500)
func GetCourseByCatalogName(catalogName string) (*model.Course, error) {
	var course model.Course
	result := postgres.DB.Where("catalog_course_name", strings.TrimSpace(catalogName)).First(&course)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &course, nil
}

// GetCourseListByMajorName gets the course list with the given major name (eg, CS)
func GetCourseListByMajorName(majorName string) ([]model.Course, error) {
	var courseList []model.Course
	majorName = majorName + "%"
	res := postgres.DB.Where("catalog_course_name LIKE ?", majorName).Find(&courseList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return courseList, nil
}
