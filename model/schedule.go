package model

type Schedule struct {
	Base
	UserID     int64 `gorm:"column:user_id"`
	SemesterID int64 `gorm:"column:semester_id"`
	ClassID    int64 `gorm:"column:class_id"`
}
