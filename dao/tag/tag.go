package tag

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
)

func GetTagList() ([]model.Tag, error) {
	var tags []model.Tag
	res := postgres.DB.Find(&tags)

	if res.Error != nil {
		return nil, res.Error
	}

	return tags, nil
}

func GetTagListByCourseID(courseID int32) ([]model.Tag, error) {
	var tags []model.Tag
	res := postgres.DB.Where("course_id = ?", courseID).Find(&tags)
	if res.Error != nil {
		return nil, res.Error
	}
	return tags, nil
}