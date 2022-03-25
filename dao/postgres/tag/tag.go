package tag

import (
	"fmt"
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

func GetTagByID(id int32) (*model.Tag, error) {
	var tag model.Tag
	res := postgres.DB.Find(&tag)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &tag, nil
}

func GetTagListByCourseID(courseID int64) ([]model.Tag, error) {
	var tags []model.Tag
	res := postgres.DB.Where("course_id = ?", courseID).Find(&tags)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return tags, nil
}

func CreateTagByCourseID(courseID int64, tagContent string, tagType int32) error {
	// if the same tag for the same course already exist, just upvote for the tag
	var existingTag model.Tag
	res := postgres.DB.Where("course_id = ? and name = ?", courseID, tagContent).Find(&existingTag)
	if res.RowsAffected != 0 {
		existingTag.VoteCount += 1
		res = postgres.DB.Model(&existingTag).Update("vote_count", existingTag.VoteCount)
		if res.Error != nil || res.RowsAffected == 0 {
			fmt.Println(res.Error.Error())
			return shared.InternalErr{}
		}
		return nil
	}

	tag := &model.Tag{
		Name:      tagContent,
		CourseID:  courseID,
		VoteCount: 1,
		Type: tagType,
	}
	res = postgres.DB.Create(&tag)
	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}

func VoteForTagByID(tagID int32, userID int64) error {
	userVotedTag := &model.UserVotedTag{
		UserID: userID,
		TagID: tagID,
	}
	res := postgres.DB.Create(&userVotedTag)
	if res.Error != nil {
		return shared.InternalErr{}
	}

	var tag model.Tag
	res = postgres.DB.Where("id = ?", tagID).Find(&tag)
	if res.Error != nil || res.RowsAffected == 0 {
		return shared.ContentNotFoundErr{}
	}

	tag.VoteCount += 1
	res = postgres.DB.Update("vote_count", tag.VoteCount)

	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}

	return nil
}
