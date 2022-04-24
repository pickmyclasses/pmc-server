package dao

import (
	"errors"
	"time"

	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"gorm.io/gorm"
)

func GetCourseOverallRating(courseID int64) (*model.OverAllRating, error) {
	var rating model.OverAllRating
	result := postgres.DB.Where("course_id = ?", courseID).First(&rating)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newRating, err := CreateCourseRating(courseID)
			if err != nil {
				return nil, shared.ContentNotFoundErr{}
			}
			rating = *newRating
		} else {
			return nil, shared.InternalErr{}
		}
	}

	return &rating, nil
}

func CreateCourseRating(CourseID int64) (*model.OverAllRating, error) {
	rating := &model.OverAllRating{
		CourseID:         CourseID,
		OverAllRating:    0,
		TotalRatingCount: 0,
	}
	result := postgres.DB.Create(&rating)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, shared.InternalErr{}
	}
	return rating, nil
}

func UpdateCourseRating(rating model.OverAllRating) error {
	res := postgres.DB.Model(&rating).
		Updates(map[string]interface{}{"over_all_rating": rating.OverAllRating, "total_rating_count": rating.TotalRatingCount})

	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}

func GetReviewTotalByCourseID(courseID int) (int64, error) {
	var total int64
	res := postgres.DB.Model(&model.Review{}).Where("course_id = ?", courseID).Count(&total)
	if res.Error != nil {
		return -1, shared.InternalErr{}
	}
	return total, nil
}

func GetPaginatedReviewsByCourseID(courseID, pn, pSize int) ([]model.Review, error) {
	var reviewList []model.Review

	res := postgres.DB.Scopes(Paginate(pn, pSize)).Where("course_id = ?", courseID).Find(&reviewList)
	if res.Error != nil {
		return make([]model.Review, 0), shared.InternalErr{}
	}
	if res.RowsAffected == 0 {
		return make([]model.Review, 0), nil
	}

	return reviewList, nil
}

func GetReviewListByCourseID(courseID int64) ([]model.Review, error) {
	var reviewList []model.Review
	res := postgres.DB.Where("course_id = ?", courseID).Find(&reviewList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return reviewList, nil
}

func GetReviewByID(reviewID int) (*model.Review, error) {
	var review model.Review
	result := postgres.DB.Where("id = ?", reviewID).First(&review)
	if result.RowsAffected == 0 {
		return nil, shared.ContentNotFoundErr{}
	}
	return &review, nil
}

func CreateCourseReview(review model.Review) error {
	res := postgres.DB.Create(&review)
	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}

func UpdateCourseReview(review model.ReviewParams) error {
	var currentReview model.Review
	res := postgres.DB.Where("course_id = ?, user_id = ?", review.CourseID, review.UserID).Find(&currentReview)
	if res.RowsAffected == 0 || res.Error != nil {
		return shared.InternalErr{}
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
		return shared.InternalErr{}
	}
	return nil
}

func GetReviewOfUserForACourse(userID, courseID int64) (*model.Review, error) {
	var review model.Review
	res := postgres.DB.Where("course_id = ? and user_id = ?", courseID, userID).First(&review)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, shared.InternalErr{}
	}
	return &review, nil
}
