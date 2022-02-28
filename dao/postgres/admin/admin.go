package dao

import (
	"pmc_server/init/postgres"
	model "pmc_server/model"
)

func FindAllUsers() ([]model.User, error) {
	var users []model.User
	resp := postgres.DB.Find(&users)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return users, nil
}
