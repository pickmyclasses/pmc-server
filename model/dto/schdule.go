package dto

import "pmc_server/model"

type Schedule struct {
	ScheduledClassList []ClassInfo `json:"scheduledClasses"`
	CustomEvents []CustomEvent `json:"customEvents"`
}

type ClassInfo struct {
	ClassData  model.Class `json:"classData"`
	CourseInfo Course  `json:"course"`
}

type CustomEvent struct {
	ID int32 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Color string `json:"color"`
	Days []int32 `json:"days"`
	StartTime int32 `json:"startTime"`
	EndTime int32 `json:"endTime"`
}
