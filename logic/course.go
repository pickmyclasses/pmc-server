package logic

import (
	"errors"
	dao "pmc_server/dao/course"
	model "pmc_server/model"
	"strconv"
)

func GetCourseList(pn, pSize int) ([]model.Course, int64) {
	return dao.GetCourses(pn, pSize)
}

func GetCourseInfo(id string) (*model.Course, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("provided ID is invalid")
	}
	return dao.GetCourseByID(idInt)
}

func GetClassListByCourseID(id string, pn, pSize int) (*[]model.Class, int64, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, 0, errors.New("provided ID is invalid")
	}
	classList, total := dao.GetClassListByCourseID(idInt, pn, pSize)
	return classList, total, nil
}
