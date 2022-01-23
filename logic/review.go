package logic

import (
	"errors"
	"strconv"

	. "pmc_server/consts"
	dao "pmc_server/dao/review"
	model "pmc_server/model"
)

func GetCourseReviewList(pn, pSize int, courseID string) (*[]model.Review, int64, error) {
	idInt, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, 0, errors.New(BAD_ID_ERR)
	}
	return dao.GetReviewsByCourseID(pn, pSize, idInt)
}

func GetReviewByID(reviewID string) (*model.Review, error) {
	idInt, err := strconv.Atoi(reviewID)
	if err != nil {
		return nil, errors.New(BAD_ID_ERR)
	}
	return dao.GetReviewByID(idInt)
}
