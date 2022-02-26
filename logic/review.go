package logic

import (
	"errors"
	"fmt"
	"strconv"

	. "pmc_server/consts"
	reviewDao "pmc_server/dao/review"
	userDao "pmc_server/dao/user"
	"pmc_server/model"
	"pmc_server/model/dto"
)

func GetCourseReviewList(pn, pSize int, courseID string) (*dto.ReviewList, error) {
	idInt, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, errors.New(BAD_ID_ERR)
	}

	rating, err := reviewDao.GetCourseReviewOverallRating(idInt)
	if err != nil {
		return nil, err
	}

	total, err := reviewDao.GetReviewTotalByCourseID(idInt)
	if err != nil {
		return nil, err
	}

	reviewList, err := reviewDao.GetReviewsByCourseID(idInt, pn, pSize)
	if err != nil {
		return nil, err
	}

	reviewRsp := &dto.ReviewList{
		Total: total,
		Reviews: make([]dto.Review, 0),
		OverallRating: rating,
	}

	for _, review := range reviewList {
		user, err := userDao.GetUserByID(review.UserID)
		if err != nil {
			return nil, err
		}

		reviewDto := dto.Review{
			ID: review.ID,
			CreatedAt: review.CreatedAt,
			Rating: review.Rating,
			Anonymous: review.Anonymous,
			Recommended: review.Recommended,
			Pros: review.Pros,
			Cons: review.Cons,
			Comment: review.Comment,
			CourseID: review.CourseID,
			UserID: review.UserID,
			Username: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		}

		reviewRsp.Reviews = append(reviewRsp.Reviews, reviewDto)
	}

	return reviewRsp, nil

}

func GetReviewByID(reviewID string) (*model.Review, error) {
	idInt, err := strconv.Atoi(reviewID)
	if err != nil {
		return nil, errors.New(BAD_ID_ERR)
	}
	return reviewDao.GetReviewByID(idInt)
}

func PostCourseReview(review dto.Review) error {
	return reviewDao.PostCourseReview(review)
}

func UpdateCourseReview(review model.ReviewParams) error {
	return reviewDao.UpdateCourseReview(review)
}
