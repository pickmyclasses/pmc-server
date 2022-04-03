package model

type Major struct {
	Base
	CollegeID int32  `json:"collegeID"`
	Name      string `json:"name"`
}
