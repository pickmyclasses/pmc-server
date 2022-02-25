package professor

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
)

func GetProfessorList() ([]model.Professor, error) {
	var professors []model.Professor
	res := postgres.DB.Find(&professors)
	if res.Error != nil {
		return nil, res.Error
	}
	return professors, nil
}
