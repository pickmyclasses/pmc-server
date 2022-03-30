package model

type CustomEvent struct {
	Base
	Title string `gorm:"title"`
	Description string `gorm:"description"`
	Color string `gorm:"color"`
	Days []int32 `gorm:"days"`
	StartTime int32 `gorm:"start_time"`
	EndTime int32 `gorm:"end_time"`
	UserID int64 `gorm:"user_id"`
	SemesterID int64 `gorm:"semester_id"`
}
