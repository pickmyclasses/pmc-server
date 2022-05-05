package model

type AssociatedTag struct {
	Base
	CourseID int64   `gorm:"column:course_id"`
	Content  string  `gorm:"column:content"`
	Weight   float32 `gorm:"column:weight"`
}
