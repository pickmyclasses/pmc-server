package model

type Bookmark struct {
	Base
	UserID   int64 `gorm:"column:user_id"`
	CourseID int64 `gorm:"column:marked_course_id"`
}
