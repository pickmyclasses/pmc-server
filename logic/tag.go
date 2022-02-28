package logic

import (
	"errors"
	courseDao "pmc_server/dao/postgres/course"
	dao "pmc_server/dao/postgres/tag"
	"pmc_server/model"
)

func GetTagList() ([]model.Tag, error) {
	return dao.GetTagList()
}

func GetTagOfCourse(courseID int64) ([]model.Tag, error) {
	return dao.GetTagListByCourseID(courseID)
}

func CreateTagByCourseID(tagInfo model.CreateTagParam) error {
	course, err := courseDao.GetCourseByID(int(tagInfo.CourseID))
	if err != nil {
		return err
	}
	if course == nil {
		return errors.New("no course found")
	}

	if len(tagInfo.Content) > 10 {
		return errors.New("tag length exceeded")
	}

	err = dao.CreateTagByCourseID(tagInfo.CourseID, tagInfo.Content)
	if err != nil {
		return err
	}
	return nil
}

func VoteTag(courseID int64, tagID int32, userID int64, ifUpvote bool) (*model.Tag, error) {
	return nil, nil
}
