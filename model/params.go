package model

type RegisterParams struct {
	Email      string `json:"email"`       // User Email
	FirstName  string `json:"first_name"`  // User first name
	LastName   string `json:"last_name"`   // User Last name
	College    string `json:"college"`     // User college, this is restricted to UofU now
	Password   string `json:"password"`    // User provided password, will be encrypted
	RePassword string `json:"re_password"` // User reentered password, to make sure they match
}

type LoginParams struct {
	Email    string `json:"email" binding:"required"`    // User Email
	Password string `json:"password" binding:"required"` // User password
}

type CourseParams struct {
	ID string `uri:"id" binding:"required"` // Course ID
}

type CourseFilterParams struct {
	// TODO: fix this with actual professor entities
	TaughtProfessor       []string `json:"taught_professor"`           // Filter courses with professor names
	Keyword               string   `json:"keyword"`                    // Keyword user inputs, this links to the name/catalog name/subject/tag of the course
	MinCredit             float32  `json:"credit"`                     // Filter courses with given credit
	MaxCredit             float32  `json:"max_credit"`                 //Filter cores with the given max credit
	OfferedOnline         bool     `json:"is_online"`                  // Filter courses that's online
	OfferedOffline        bool     `json:"is_offline"`                 // Filter courses that's not online (in person)
	IsHonor               bool     `json:"is_honor"`                   // Filter courses that's honor courses
	Weekday               []int8   `json:"weekday"`                    // Filter courses that's in the specific weekdays
	StartTime             float32  `json:"start_time"`                 // Filter courses that starts no earlier than the start time
	EndTime               float32  `json:"end_time"`                   // Filter courses that ends no later than the start time
	MinRating             int8     `json:"min_rating"`                 // Filter courses that has no lower rating than the give min rating
	RankByRatingHighToLow bool     `json:"rank_by_rating_high_to_low"` // Rank the courses by the given rating low to high
	RankByRatingLowToHigh bool     `json:"rank_by_rating_low_to_high"` // Rank the courses by the given rating low to high
	PageNumber            int      `json:"page_number"`                // The current page of the search result, default 0
	PageSize              int      `json:"page_size"`                  // The current page size of the search result, default 10
}

type ClassParams struct {
	ID string `uri:"id" binding:"required"` // Class ID
}

type ReviewParams struct {
	CourseID    int64   `json:"course_id" binding:"required"`    // Course ID
	UserID      int64   `json:"user_id" binding:"required"`      // User ID
	Pros        string  `json:"pros" binding:"required"`         // The pros of the course given by the user
	Cons        string  `json:"cons" binding:"required"`         // The cons of the course given by the user
	Comment     string  `json:"comments" binding:"required"`     // The detailed comment on the review given by the user
	Rating      float32 `json:"rating" binding:"required"`       // The rating of the course given by the user
	IsAnonymous bool    `json:"is_anonymous" binding:"required"` // Is the user posting this review anonymously
	Recommended bool    `json:"recommended" binding:"required"`  // Is the course recommended by the user
}

type PostScheduleParams struct {
	ClassID    int64 `json:"class_id"`    // Class ID
	SemesterID int64 `json:"semester_id"` // Semester ID
	UserID     int64 `json:"user_id"`  // Student ID
}

type GetScheduleParams struct {
	UserID int64 `form:"user_id"` // Student ID
}

type DeleteScheduleParams struct {
	UserID     int64 `json:"user_id"`     // Student ID
	ClassID    int64 `json:"class_id"`    // Class ID
	SemesterID int64 `json:"semester_id"` // Semester ID
}

type GetTagParams struct {
	CourseID int32 `json:"course_id"` // Course ID
}
