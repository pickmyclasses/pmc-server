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
	IsHot              bool     `json:"is_hot"`               // Should the courses be hot
	IsOnline           bool     `json:"is_online"`            // Should the courses be online
	IsRankRating       bool     `json:"is_rank_rating"`       // Should the courses be ranked as rating high to low
	IsMajorRequirement bool     `json:"is_major_requirement"` // Should the courses be the major requirements
	WeekdaySelected    []uint8  `json:"weekday_select"`       // Should the courses only be taught in a certain day
	TagsSelected       []string `json:"tags_selected"`        // Should the courses be associated to certain tags
	// TODO: change this to professor object
	ProfessorSelected []string `json:"professor_selected"` // Should the courses be taught be certain professors
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
	Recommended bool    `json:"recommended" binding:"required"`  // Is teh course recommended by the user
}

type PostScheduleParams struct {
	ClassID    int64 `json:"class_id"`    // Class ID
	SemesterID int64 `json:"semester_id"` // Semester ID
	UserID     int64 `json:"student_id"`  // Student ID
}

type GetScheduleParams struct {
	UserID int64 `json:"user_id"` // Student ID
}

type DeleteScheduleParams struct {
	UserID     int64 `json:"user_id"`     // Student ID
	ClassID    int64 `json:"class_id"`    // Class ID
	SemesterID int64 `json:"semester_id"` // Semester ID
}
