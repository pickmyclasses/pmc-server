package model

type UserCourseHistory struct {
	Base
	UserID int64 `gorm:"column:user_id"`
	CourseID int64 `gorm:"column:course_id"`
	ClassID int64 `gorm:"column:class_id"`
	SemesterID int32 `gorm:"column:semester_id"`
}
