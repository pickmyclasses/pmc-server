package model

type Semester struct {
	Base
	CollegeID int32 `gorm:"column:college_id"`
	Season string `gorm:"column:season"`
	Year int32 `gorm:"column:year"`
}
