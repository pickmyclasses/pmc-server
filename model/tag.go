package model

type Tag struct {
	Base
	Name      string `gorm:"column:name" json:"name"`
	VoteCount int32  `gorm:"column:vote_count" json:"voteCount"`
	CourseID  int64  `gorm:"column:course_id" json:"courseID"`
}
