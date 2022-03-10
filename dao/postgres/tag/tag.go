package tag

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetTagList() ([]model.Tag, error) {
	var tags []model.Tag
	res := postgres.DB.Find(&tags)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}

	return tags, nil
}

func GetTagListByCourseID(courseID int64) ([]model.Tag, error) {
	var tags []model.Tag
	res := postgres.DB.Where("course_id = ?", courseID).Find(&tags)
	if res.Error != nil {
		return nil, shared.InternalErr{}
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
		return shared.InternalErr{}
	}
	return nil
}
