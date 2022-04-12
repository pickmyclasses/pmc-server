package logic

import (
	courseDao "pmc_server/dao/postgres/course"
	dao "pmc_server/dao/postgres/tag"
	"pmc_server/model"
	"pmc_server/shared"
)

type tagDto struct {
	name string `json:"name"`
	id   int64  `json:"id"`
}

func GetTagList() ([]tagDto, error) {
	tagList, err := dao.GetTagList()
	if err != nil {
		return nil, err
	}
	tagDtoList := make([]tagDto, 0)
	seen := make(map[int64]bool, 0)
	for _, tag := range tagList {
		if v, exist := seen[tag.ID]; exist && v {
			continue
		}
		dto := tagDto{
			name: tag.Name,
			id:   tag.ID,
		}
		tagDtoList = append(tagDtoList, dto)
		seen[tag.ID] = true
	}
	return tagDtoList, nil
}

func GetTagOfCourse(courseID int64) ([]model.Tag, error) {
	return dao.GetTagListByCourseID(courseID)
}

func CreateTagByCourseID(content string, tagType int32, courseID int64) error {
	course, err := courseDao.GetCourseByID(int(courseID))
	if err != nil {
		return err
	}
	if course == nil {
		return shared.ContentNotFoundErr{}
	}

	if len(content) > 20 {
		content = content[:20]
	}

	err = dao.CreateTagByCourseID(courseID, content, tagType)
	if err != nil {
		return err
	}
	return nil
}

func VoteTag(tagID int32, userID int64) error {
	tag, err := dao.GetTagByID(tagID)
	if err != nil {
		return err
	}
	if tag == nil {
		return shared.ContentNotFoundErr{}
	}

	err = dao.VoteForTagByID(tagID, userID)
	if err != nil {
		return err
	}

	return nil
}
