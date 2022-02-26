package dto

import "pmc_server/model"

type Schedule struct {
	ScheduledClassList []ClassInfo `json:"scheduled_class_list"`
}

type ClassInfo struct {
	Rating	float32 `json:"rating"`
	ClassData model.Class 	`json:"class_data"`
	CourseData model.Course `json:"course_data"`
}

