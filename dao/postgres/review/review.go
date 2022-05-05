package dao

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
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
	res := postgres.DB.Model(&rating).Where("course_id = ?", rating.CourseID).
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

	res := postgres.DB.Scopes(shared.Paginate(pn, pSize)).Where("course_id = ?", courseID).Find(&reviewList)
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

func UpdateCourseReview(review model.Review) error {
	var currentReview model.Review
	res := postgres.DB.Where("course_id = ? and user_id = ?", review.CourseID, review.UserID).Find(&currentReview)
	if res.RowsAffected == 0 || res.Error != nil {
		return shared.InternalErr{}
	}

	currentReview.CreatedAt = time.Now()
	currentReview.Anonymous = review.Anonymous
	currentReview.Rating = review.Rating
	currentReview.Comment = review.Comment
	currentReview.Recommended = review.Recommended
	currentReview.Tags = review.Tags
	currentReview.HomeworkHeavy = review.HomeworkHeavy
	currentReview.ExamHeavy = review.ExamHeavy
	currentReview.HourSpent = review.HourSpent
	currentReview.GradeReceived = review.GradeReceived

	res = postgres.DB.Save(&currentReview)
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

func CheckReviewExist(userID, courseID int64) (bool, error) {
	var review model.Review
	res := postgres.DB.Where("course_id = ? and user_id = ?", courseID, userID).First(&review)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, shared.InternalErr{}
	}
	return true, nil
}

func UpdateReviewVotes(courseID, userID int64, voterID int64, isUpvote bool) error {
	var review model.Review
	res := postgres.DB.Where("course_id  = ? and user_id = ?", courseID, userID).First(&review)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return shared.NoPreviousRecordErr{}
		}
		return shared.InternalErr{}
	}

	alreadyVoted := true
	var userVotedReview model.UserVotedReview
	res = postgres.DB.Where("user_id = ? and reviewer_id = ? and course_id = ?", voterID, userID, courseID).First(&userVotedReview)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			alreadyVoted = false
		} else {
			return shared.InternalErr{}
		}
	}

	if alreadyVoted {
		if userVotedReview.IsUpvote == isUpvote {
			return shared.ResourceConflictErr{}
		}
		// if already liked/disliked, update the original
		userVotedReview.IsUpvote = isUpvote
		res = postgres.DB.Save(&userVotedReview)

		// update the like, disliked count for review
		if isUpvote == false {
			review.LikeCount--
			review.DislikeCount++
		} else {
			review.LikeCount++
			review.DislikeCount--
		}

		res = postgres.DB.Save(&review)
		if res.Error != nil {
			return shared.InternalErr{}
		}

		return nil
	}

	// update the vote count
	if isUpvote {
		review.LikeCount += 1
	} else {
		review.DislikeCount += 1
	}

	res = postgres.DB.Save(&review)
	if res.Error != nil {
		return shared.InternalErr{}
	}

	// update the user voted history
	userVoted := model.UserVotedReview{
		UserID:     voterID,
		ReviewerID: userID,
		CourseID:   courseID,
		IsUpvote:   isUpvote,
	}
	res = postgres.DB.Create(&userVoted)
	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{}
	}
	return nil
}

func GetUserVotedReviews(userID int64, courseID int64) ([]model.UserVotedReview, error) {
	var reviewHistory []model.UserVotedReview
	res := postgres.DB.Where("user_id = ? and course_id = ?", userID, courseID).Find(&reviewHistory)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []model.UserVotedReview{}, nil
		}
		return nil, shared.InternalErr{}
	}
	return reviewHistory, nil
}
