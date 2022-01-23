package dto

type Class struct {
	ID                  int64  `json:"id"`
	Semester            string `json:"semester"`
	Year                string `json:"year"`
	Session             string `json:"session"`
	WaitList            string `json:"wait_list"`
	OfferDate           string `json:"offer_date"`
	StartTime           string `json:"start_time"`
	EndTime             string `json:"end_time"`
	Location            string `json:"location"`
	RecommendationScore int    `json:"recommendation_score"`
	Type                string `json:"type"`
	Number              string `json:"number"`
	Component           string `json:"component"`
	Unit                string `json:"unit"`
	SeatAvailable       int32  `json:"seat_available"`
	Note                string `json:"note"`
	Instructors         string `json:"instructors"`
	CourseTitle         string `json:"course_title"`
	CourseCatalogName   string `json:"course_catalog_name"`
	CourseID            int64  `json:"course_id"`
}
