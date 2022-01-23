package model

type Class struct {
	Base
	Semester            string  `gorm:"column:semester"`
	Year                string  `gorm:"column:year"`
	Session             string  `gorm:"column:session"`
	WaitList            string  `gorm:"column:wait_list"`
	OfferDate           string  `gorm:"column:offer_date"`
	StartTime           string  `gorm:"column:start_time"`
	EndTime             string  `gorm:"column:end_time"`
	Location            string  `gorm:"column:location"`
	RecommendationScore int     `gorm:"column:recommendation_score"`
	Type                string  `gorm:"column:type"`
	Number              string  `gorm:"column:number"`
	Component           string  `gorm:"column:component"`
	Unit                string  `gorm:"column:unit"`
	SeatAvailable       int32   `gorm:"column:seat_available"`
	Note                string  `gorm:"column:notes"`
	Instructors         string  `gorm:"column:instructors"`
	CourseTitle         string  `gorm:"column:course_title"`
	CourseCatalogName   string  `gorm:"column:course_catalog_name"`
	CourseID            int64   `gorm:"column:course_id"`
	Course              *Course `gorm:"foreignKey:CourseID"`
}
