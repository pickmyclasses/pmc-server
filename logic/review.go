package logic

import (
	"fmt"
	"strconv"

	reviewDao "pmc_server/dao/postgres/review"
	userDao "pmc_server/dao/postgres/user"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
)

func GetCourseReviewList(pn, pSize int, courseID string) (*dto.ReviewList, error) {
	idInt, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, shared.MalformedIDErr{}
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
		var username string
		userInfo, err := userDao.GetUserByID(review.UserID)
		if err != nil || review.Anonymous {
			username = "Anonymous student"
		} else {
			username = fmt.Sprintf("%s %s", userInfo.FirstName, userInfo.LastName)
		}
		reviewDto := dto.Review{
			Rating:      review.Rating,
			Anonymous:   review.Anonymous,
			Recommended: review.Recommended,
			Pros:        review.Pros,
			Cons:        review.Cons,
			Comment:     review.Comment,
			CourseID:    int64(idInt),
			UserID:      review.UserID,
			Username:    username,
			CreatedAt:   review.CreatedAt,
		}

		reviewRsp.Reviews = append(reviewRsp.Reviews, reviewDto)
	}

	return reviewRsp, nil

}

func GetReviewByID(reviewID string) (*model.Review, error) {
	idInt, err := strconv.Atoi(reviewID)
	if err != nil {
		return nil, shared.MalformedIDErr{}
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
		((rating.OverAllRating * float32(rating.TotalRatingCount)) + review.Rating) /
			(float32(rating.TotalRatingCount) + 1)

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
