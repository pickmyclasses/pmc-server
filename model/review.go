package model

type Review struct {
	Base
	Rating       float32 `gorm:"column:rating"`
	Anonymous    bool    `gorm:"column:anonymous"`
	Recommended  bool    `gorm:"column:recommended"`
	Pros         string  `gorm:"column:pros"`
	Cons         string  `gorm:"column:cons"`
	Comment      string  `gorm:"column:comment"`
	CourseID     int64   `gorm:"column:course_id"`
	UserID       int64   `gorm:"column:user_id"`
	LikeCount    int32   `gorm:"column:like_count"`
	DislikeCount int32   `gorm:"column:dislike_count"`
}
