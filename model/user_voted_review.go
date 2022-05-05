package model

type UserVotedReview struct {
	Base
	UserID     int64 `gorm:"user_id"`
	ReviewerID int64 `gorm:"reviewer_id"`
	CourseID   int64 `gorm:"course_id"`
	IsUpvote   bool  `gorm:"is_upvote"`
}
