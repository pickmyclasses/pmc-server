package model

type OverAllRating struct {
	Base
	CourseID         int64   `gorm:"column:course_id"`
	OverAllRating    float32 `gorm:"column:over_all_rating"`
	TotalRatingCount int32   `gorm:"column:total_rating_count"`
}
