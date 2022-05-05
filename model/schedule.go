package model

type Schedule struct {
	Base
	UserID     int64 `gorm:"column:user_id" json:"userID"`
	SemesterID int64 `gorm:"column:semester_id" json:"semesterID"`
	ClassID    int64 `gorm:"column:class_id" json:"classID"`
}
