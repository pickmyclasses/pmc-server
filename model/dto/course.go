package dto

import "pmc_server/model"

type Course struct {
	CourseID           int64         `json:"id"`
	IsHonor            bool          `json:"isHonor"`
	FixedCredit        bool          `json:"fixedCredit"`
	DesignationCatalog string        `json:"designationCatalog"`
	Description        string        `json:"description"`
	Prerequisites      string        `json:"prerequisites"`
	Title              string        `json:"title"`
	CatalogCourseName  string        `json:"catalogCourseName"`
	Component          string        `json:"component"`
	MaxCredit          float64       `json:"maxCredit"`
	MinCredit          float64       `json:"minCredit"`
	Classes            []model.Class `json:"classes"`       // The classes (sessions) of the given course
	OverallRating      float32       `json:"overallRating"` // The overall rating of the class
	Tags               []model.Tag   `json:"tags"`
}
