package logic

import (
	"errors"
	"pmc_server/model/dto"
	"strconv"

	. "pmc_server/consts"
	dao "pmc_server/dao/review"
	model "pmc_server/model"
)

func GetCourseReviewList(pn, pSize int, courseID string) (*[]dto.Review, int64, error) {
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

func PostCourseReview(review dto.Review) error {
	return dao.PostCourseReview(review)
}

func UpdateCourseReview(review model.ReviewParams) error {
	return dao.UpdateCourseReview(review)
}
