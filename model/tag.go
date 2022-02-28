package model

type Tag struct {
	Base
	Name      string `gorm:"column:name"`
	VoteCount int32  `gorm:"column:vote_count"`
	CourseID  int64  `gorm:"column:course_id"`
}
