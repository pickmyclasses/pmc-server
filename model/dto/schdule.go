package dto

import "pmc_server/model"

type Schedule struct {
	ScheduledClassList []ClassInfo `json:"scheduledClassList"`
}

type ClassInfo struct {
	ClassData  model.Class  `json:"classData"`
	CourseInfo CourseInfo `json:"courseData"`
}

type CourseInfo struct {
	OverallRating float32 `json:"overallRating"`
	CourseData model.Course `json:"courseInfo"`
	CourseTags []model.Tag `json:"tags"`
}
