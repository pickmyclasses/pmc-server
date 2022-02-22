package model

type Tag struct {
	Base
	Name      string `gorm:"type:varchar(10);not null;index:idx_name"`
	Category  string `gorm:"type:varchar(20);not null"`
	VoteCount int32  `gorm:"column:vote_count;default:0;not null"`
	CourseID  int64  `gorm:"column:course_id"`
}
