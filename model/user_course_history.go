package model

type UserCourseHistory struct {
	Base
	UserID     int64 `gorm:"column:user_id"`
	CourseID   int64 `gorm:"column:course_id"`
	ProfessorName string `gorm:"column:professor_name""`
	SemesterID int32 `gorm:"column:semester_id"`
}
