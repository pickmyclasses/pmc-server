package professor

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetProfessorList() ([]model.Professor, error) {
	var professors []model.Professor
	res := postgres.DB.Find(&professors)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return professors, nil
}
