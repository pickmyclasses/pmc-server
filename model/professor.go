package model

type Professor struct {
	Base
	Name      string `gorm:"column:name"`
	CollegeID int64  `gorm:"column:college_id"`
}
