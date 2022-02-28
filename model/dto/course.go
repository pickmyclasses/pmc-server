package dto

import "pmc_server/model"

type Course struct {
	Course        *model.Course `json:"course"`         // The course info of the given course ID
	Classes       []model.Class `json:"classes"`        // The classes (sessions) of the given course
	OverallRating float32       `json:"overallRating"` // The overall rating of the class
}
