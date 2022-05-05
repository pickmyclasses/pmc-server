package logic

import (
	courseDao "pmc_server/dao/postgres/course"
	dao "pmc_server/dao/postgres/tag"
	"pmc_server/model"
	"pmc_server/shared"
	"strings"
)

type tagDto struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

func GetTagList() ([]tagDto, error) {
	tagList, err := dao.GetTagList()
	if err != nil {
		return nil, err
	}
	tagDtoList := make([]tagDto, 0)
	seen := make(map[string]int64, 0)
	for _, tag := range tagList {
		if _, exist := seen[strings.TrimSpace(strings.ToLower(tag.Name))]; exist {
			continue
		}
		dto := tagDto{
			Name: strings.TrimSpace(tag.Name),
			ID:   tag.ID,
		}
		tagDtoList = append(tagDtoList, dto)
		seen[strings.TrimSpace(strings.ToLower(tag.Name))] = tag.ID
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
