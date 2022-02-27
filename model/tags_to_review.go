package model

type TagsToReview struct {
	Base
	ReviewID int64 `gorm:"review_id"`
	UserID   int64 `gorm:"user_id"`
}
