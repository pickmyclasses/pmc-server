package dao

import (
	"errors"

	. "pmc_server/consts"
	"pmc_server/init/postgres"
	model "pmc_server/models"
	"pmc_server/utils"
)

func GetReviewsByCourseID(pn, pSize, courseID int) (*[]model.Review, int64, error) {
	var reviews []model.Review
	result := postgres.DB.Where(&model.Review{CourseID: int64(courseID)}).Find(&reviews)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	total := result.RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Find(&reviews)

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
