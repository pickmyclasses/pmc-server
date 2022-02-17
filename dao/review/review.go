package dao

import (
	"errors"
	"pmc_server/model/dto"
	"time"

	. "pmc_server/consts"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/utils"
)

func GetReviewsByCourseID(pn, pSize, courseID int) (*[]dto.Review, int64, error) {
	var reviews []model.Review
	result := postgres.DB.Where("course_id = ?", courseID).Find(&reviews)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var reviewDtos []dto.Review
	for _, review := range reviews {
		var user model.User
		_ = postgres.DB.Where("id = ?", review.UserID).Find(&user)
		reviewDto := &dto.Review{
			ID:          review.ID,
			CreatedAt:   review.CreatedAt,
			Rating:      review.Rating,
			Anonymous:   review.Anonymous,
			Recommended: review.Recommended,
			Pros:        review.Pros,
			Cons:        review.Cons,
			Comment:     review.Comment,
			CourseID:    review.CourseID,
			UserID:      review.UserID,
			Username:    user.FirstName + " " + user.LastName,
		}
		reviewDtos = append(reviewDtos, *reviewDto)
	}

	total := result.RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Where("course_id = ?", courseID).Find(&reviews)

	return &reviewDtos, total, nil
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

func UpdateCourseReview(review model.ReviewParams) error {
	var currentReview model.Review
	res := postgres.DB.Where("course_id = ?, user_id = ?", review.CourseID, review.UserID).Find(&currentReview)
	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("no review of given info found")
	}

	currentReview.CreatedAt = time.Now()
	currentReview.Anonymous = review.IsAnonymous
	currentReview.Rating = review.Rating
	currentReview.Comment = review.Comment
	currentReview.Pros = review.Pros
	currentReview.Cons = review.Cons
	currentReview.Recommended = review.Recommended
	res = postgres.DB.Updates(&currentReview)
	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("internal error when updating the review")
	}
	return nil
}
