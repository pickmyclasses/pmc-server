package dto

import "pmc_server/model"

type Course struct {
	Course  *model.Course `json:"course"`
	Classes []model.Class `json:"classes"`
}
