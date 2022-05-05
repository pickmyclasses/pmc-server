package dto

type Class struct {
	ID                  int64  `json:"id"`
	Semester            string `json:"semester"`
	Year                string `json:"year"`
	Session             string `json:"session"`
	WaitList            string `json:"waitList"`
	OfferDate           string `json:"offerDate"`
	StartTime           string `json:"startTime"`
	EndTime             string `json:"endTime"`
	Location            string `json:"location"`
	RecommendationScore int    `json:"recommendationScore"`
	Type                string `json:"type"`
	Number              string `json:"number"`
	Component           string `json:"component"`
	Unit                string `json:"unit"`
	SeatAvailable       int32  `json:"seatAvailable"`
	Note                string `json:"note"`
	Instructors         string `json:"instructors"`
	CourseTitle         string `json:"courseTitle"`
	CourseCatalogName   string `json:"courseCatalogName"`
	CourseID            int64  `json:"courseId"`
}
