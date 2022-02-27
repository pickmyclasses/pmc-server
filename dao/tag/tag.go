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

func GetTagListByCourseID(courseID int64) ([]model.Tag, error) {
	var tags []model.Tag
	res := postgres.DB.Where("course_id = ?", courseID).Find(&tags)
	if res.Error != nil {
		return nil, res.Error
	}
	return tags, nil
}

func CreateTagByCourseID(courseID int64, tagContent string) error {
	tag := &model.Tag{
		Name:      tagContent,
		CourseID:  courseID,
		VoteCount: 1,
	}
	res := postgres.DB.Create(&tag)
	if res.Error != nil || res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}
