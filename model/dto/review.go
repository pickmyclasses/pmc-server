package dto

type Review struct {
	Rating      float32 `json:"rating" binding:"required"`      // The giving rating of the review
	Anonymous   bool    `json:"anonymous" binding:"required"`   // Is the review posted anonymously
	Recommended bool    `json:"recommended" binding:"required"` // Is the review recommend
	Pros        string  `json:"pros" binding:"required"`        // The written text under pros
	Cons        string  `json:"cons" binding:"required"`        // The written text under cons
	Comment     string  `json:"comment" binding:"required"`     // The overall comment
	CourseID    int64   `json:"course_id" binding:"required"`   // The course ID of the review
	UserID      int64   `json:"user_id" binding:"required"`     // The userID who posted this review
}
