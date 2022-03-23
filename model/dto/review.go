package dto

import "time"

type ReviewList struct {
	Total         int64    `json:"total"`         // The total number of reviews of the course
	Reviews       []Review `json:"reviews"`       // The review set
	OverallRating float32  `json:"overallRating"` // The over all rating of the course
}

type Review struct {
	Rating                  float32   `json:"rating"`      // The giving rating of the review
	Anonymous               bool      `json:"anonymous"`   // Is the review posted anonymously
	Recommended             bool      `json:"recommended"` // Is the review recommend
	Pros                    string    `json:"pros" `       // The written text under pros
	Cons                    string    `json:"cons" `       // The written text under cons
	Comment                 string    `json:"comment" `    // The overall comment
	CourseID                int64     `json:"courseID"`    // The course ID of the review
	UserID                  int64     `json:"userID"`      // The userID who posted this review
	Username                string    `json:"username"`
	CreatedAt               time.Time `json:"createdAt"`
	LikedCount              int32     `json:"likedCount"`
	DislikedCount           int32     `json:"dislikedCount"`
	HourSpent int      `json:"hourSpent"`
	GradeReceived           string    `json:"gradeReceived"`
	IsExamHeavy             bool      `json:"isExamHeavy"`
	IsHomeworkHeavy         bool      `json:"isHomeworkHeavy"`
	ExtraCreditOffered      bool      `json:"extraCreditOffered"`
}
