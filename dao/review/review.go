package dao

import (
	"errors"
	"time"

	. "pmc_server/consts"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/utils"
)

func GetCourseOverallRating(courseID int64) (*model.OverAllRating, error) {
	var rating model.OverAllRating
	result := postgres.DB.Where("course_id = ?", courseID).First(&rating)
	if result.Error != nil {
		return nil, errors.New("failed to get course rating")
	}

	return &rating, nil
}

func CreateCourseRating(CourseID int64) (*model.OverAllRating, error) {
	rating := &model.OverAllRating{
		CourseID: CourseID,
		OverAllRating: 0,
		TotalRatingCount: 0,
	}
	result := postgres.DB.Create(&rating)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("create new rating record failed")
	}
	return rating, nil
}

func UpdateCourseRating(rating model.OverAllRating) error {
	res := postgres.DB.
		Where("course_id = ?", rating.CourseID).
		Updates(map[string]interface{}{"over_all_rating": rating.OverAllRating, "total_rating_count": rating.TotalRatingCount})

	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New("update course rating failed")
	}
	return nil
}

func GetReviewTotalByCourseID(courseID int) (int64, error) {
	var total int64
	res := postgres.DB.Where("course_id = ?", courseID).Count(&total)
	if res.Error != nil {
		return -1, errors.New("failed to fetch total number of review")
	}
	return total, nil
}

func GetReviewsByCourseID(courseID, pn, pSize int) ([]model.Review, error) {
	var reviewList []model.Review

	res := postgres.DB.Scopes(utils.Paginate(pn, pSize)).Where("course_id = ?", courseID).Find(&reviewList)
	if res.Error != nil {
		return nil, errors.New("failed to fetch review list")
	}
	return reviewList, nil

}

func GetReviewByID(reviewID int) (*model.Review, error) {
	var review model.Review
	result := postgres.DB.Where("id = ?", reviewID).First(&review)
	if result.RowsAffected == 0 {
		return nil, errors.New(NO_INFO_ERR)
	}
	return &review, nil
}

func CreateCourseReview(review model.Review) error {
	res := postgres.DB.Create(&review)
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
