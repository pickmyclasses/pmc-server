package dto

import "pmc_server/model"

type Schedule struct {
	ScheduledClassList []ClassInfo `json:"scheduledClassList"`
}

type ClassInfo struct {
	ClassData  model.Class `json:"classData"`
	CourseInfo Course  `json:"course"`
}
