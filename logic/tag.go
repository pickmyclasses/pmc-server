package logic

import (
	dao "pmc_server/dao/tag"
	"pmc_server/model"
)

func GetTagList() ([]model.Tag, error) {
	return dao.GetTagList()
}

func GetTagOfCourse(params model.GetTagParams) ([]model.Tag, error) {
	return dao.GetTagListByCourseID(params.CourseID)
}