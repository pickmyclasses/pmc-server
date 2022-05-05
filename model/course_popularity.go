package model

type CoursePopularity struct {
	Base
	CourseID   int64 `gorm:"course_id"`
	Popularity int32 `gorm:"popularity"`
	SemesterID int32 `gorm:"semester_id"`
}
