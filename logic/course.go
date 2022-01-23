package logic

import (
	"errors"
	"strconv"

	classDao "pmc_server/dao/class"
	dao "pmc_server/dao/course"
	"pmc_server/model"
	"pmc_server/model/dto"
)

func GetCourseList(pn, pSize int) ([]model.Course, int64) {
	return dao.GetCourses(pn, pSize)
}

func GetCourseInfo(id string) (*dto.Course, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("provided ID is invalid")
	}
	course, err := dao.GetCourseByID(idInt)
	if err != nil {
		return nil, err
	}

	classList, err := classDao.GetClassByCourseID(course.ID)
	if err != nil {
		return &dto.Course{
			Course:  course,
			Classes: nil,
		}, nil
	}
	return &dto.Course{
		Course:  course,
		Classes: *classList,
	}, nil
}

func GetClassListByCourseID(id string, pn, pSize int) (*[]model.Class, int64, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, 0, errors.New("provided ID is invalid")
	}
	classList, total := dao.GetClassListByCourseID(idInt, pn, pSize)
	return classList, total, nil
}
