package model

type RegisterParams struct {
	Email      string `json:"email" binding:"required"`       // User Email
	FirstName  string `json:"first_name" binding:"required"`  // User first name
	LastName   string `json:"last_name" binding:"required"`   // User Last name
	College    string `json:"college"`                        // User college, this is restricted to UofU now
	Password   string `json:"password" binding:"required"`    // User provided password, will be encrypted
	RePassword string `json:"re_password" binding:"required"` // User reentered password, to make sure they match
}

type LoginParams struct {
	Email    string `json:"email" binding:"required"`    // User Email
	Password string `json:"password" binding:"required"` // User password
}

type CourseParams struct {
	ID string `uri:"id" binding:"required"` // Course ID
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
