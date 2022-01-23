package model

type RegisterParams struct {
	Email      string `json:"email" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	College    string `json:"college"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

type LoginParams struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

type CourseParams struct {
	ID string `uri:"id" binding:"required"`
}

type ClassParams struct {
	ID string `uri:"id" binding:"required"`
}

type ReviewParams struct {
	CourseID int64 `json:"course_id" binding:"required"`
	UserID   int64 `json:"user_id" binding:"required"`
}
