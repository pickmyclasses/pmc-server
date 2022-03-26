package dto

import "pmc_server/model"

type Schedule struct {
	ScheduledClassList []ClassInfo `json:"scheduledClassList"`
}

type ClassInfo struct {
	Rating     float32      `json:"rating"`
	ClassData  model.Class  `json:"classData"`
	CourseData model.Course `json:"courseData"`
	CourseTags []model.Tag `json:"courseTags"`
}
