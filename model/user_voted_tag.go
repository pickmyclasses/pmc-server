package model

type UserVotedTag struct {
	Base
	TagID    int32 `gorm:"column:tag_id"`
	UserID   int64 `gorm:"column:user_id"`
	IfUpvote bool  `gorm:"column:if_upvote"`
}
