package dto

type User struct {
	ID         int64  `json:"id"`
	Token      string `json:"token"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Role       int32  `json:"role"`
	CollegeID  int32  `json:"collegeID"`
	Major      string `json:"major"`
	Emphasis   string `json:"emphasis"`
	SchoolYear string `json:"schoolYear"`
}
