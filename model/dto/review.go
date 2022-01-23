package dto

type Review struct {
	Rating      float32 `json:"rating"`      // The giving rating of the review
	Anonymous   bool    `json:"anonymous"`   // Is the review posted anonymously
	Recommended bool    `json:"recommended"` // Is the review recommend
	Pros        string  `json:"pros"`        // The written text under pros
	Cons        string  `json:"cons"`        // The written text under cons
	Comment     string  `json:"comment"`     // The overall comment
	CourseID    int64   `json:"course_id"`   // The course ID of the review
	UserID      int64   `json:"user_id"`     // The userID who posted this review
}
