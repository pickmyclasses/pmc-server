package dao

import (
	"errors"
	"pmc_server/model/dto"

	. "pmc_server/consts"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/utils"
)

func GetReviewsByCourseID(pn, pSize, courseID int) (*[]model.Review, int64, error) {
	var reviews []model.Review
	result := postgres.DB.Where("course_id = ?", courseID).Find(&reviews)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	total := result.RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Where("course_id = ?", courseID).Find(&reviews)

	return &reviews, total, nil
}

func GetReviewByID(reviewID int) (*model.Review, error) {
	var review model.Review
	result := postgres.DB.Where("id = ?", reviewID).First(&review)
	if result.RowsAffected == 0 {
		return nil, errors.New(NO_INFO_ERR)
	}
	return &review, nil
}

func PostCourseReview(review dto.Review) error {
	reviewModel := &model.Review{
		Rating:      review.Rating,
		UserID:      review.UserID,
		CourseID:    review.CourseID,
		Anonymous:   review.Anonymous,
		Recommended: review.Recommended,
		Pros:        review.Pros,
		Cons:        review.Cons,
		Comment:     review.Comment,
	}

	res := postgres.DB.Create(&reviewModel)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("create review failed")
	}
	return nil
}
