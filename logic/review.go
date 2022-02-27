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
		user, err := userDao.GetUserByID(review.UserID)
		if err != nil {
			return nil, err
		}

		reviewDto := dto.Review{
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
			Username:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
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
	// check old overall rating value
	rating, err := reviewDao.GetCourseOverallRating(review.CourseID)
	if err != nil {
		return err
	}

	// if there is no existing rating, create a new one first
	if rating == nil || rating.ID == 0 {
		rating, err = reviewDao.CreateCourseRating(review.CourseID)
		if err != nil {
			return err
		}
	}

	// recalculate the rating for the record
	rating.OverAllRating =
		((rating.OverAllRating * float32(rating.TotalRatingCount)) + review.Rating) / float32(rating.TotalRatingCount)


	reviewRec := &model.Review{
		Rating: rating.OverAllRating,
		Anonymous: review.Anonymous,
		Recommended: review.Recommended,
		Pros: review.Pros,
		Cons: review.Cons,
		Comment: review.Comment,
		CourseID: review.CourseID,
		UserID: review.UserID,
		LikeCount: 0,
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
