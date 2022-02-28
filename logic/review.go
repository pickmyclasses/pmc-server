package logic

import (
	"errors"
	reviewDao "pmc_server/dao/postgres/review"
	"strconv"

	. "pmc_server/consts"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/utils"
)

func GetCourseReviewList(pn, pSize int, courseID string) (*dto.ReviewList, error) {
	idInt, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, errors.New(BAD_ID_ERR)
	}

	rating, err := reviewDao.GetCourseOverallRating(int64(idInt))
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
		Total:         total,
		Reviews:       make([]dto.Review, 0),
		OverallRating: rating.OverAllRating,
	}

	for _, review := range reviewList {
		reviewDto := dto.Review{
			Rating:      review.Rating,
			Anonymous:   review.Anonymous,
			Recommended: review.Recommended,
			Pros:        review.Pros,
			Cons:        review.Cons,
			Comment:     review.Comment,
			CourseID:    review.CourseID,
			UserID:      review.UserID,
			CreatedAt: review.CreatedAt,
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

func PostCourseReview(review dto.Review, courseID int64) error {
	// check old overall rating value
	rating, err := reviewDao.GetCourseOverallRating(courseID)
	if err != nil {
		return err
	}

	// if there is no existing rating, create a new one first
	if rating == nil || rating.ID == 0 {
		rating, err = reviewDao.CreateCourseRating(courseID)
		if err != nil {
			return err
		}
	}

	// recalculate the rating for the record
	rating.OverAllRating =
		float32(utils.ToFixed(float64(((rating.OverAllRating*float32(rating.TotalRatingCount))+review.Rating)/
			(float32(rating.TotalRatingCount)+1)), 2))

	reviewRec := &model.Review{
		Rating:       review.Rating,
		Anonymous:    review.Anonymous,
		Recommended:  review.Recommended,
		Pros:         review.Pros,
		Cons:         review.Cons,
		Comment:      review.Comment,
		CourseID:     courseID,
		UserID:       review.UserID,
		LikeCount:    0,
		DislikeCount: 0,
	}

	err = reviewDao.CreateCourseReview(*reviewRec)
	if err != nil {
		return err
	}

	rating.TotalRatingCount += 1

	// save the new rating
	err = reviewDao.UpdateCourseRating(*rating)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCourseReview(review model.ReviewParams) error {
	return reviewDao.UpdateCourseReview(review)
}
