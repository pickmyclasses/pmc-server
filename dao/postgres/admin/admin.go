package dao

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

// FindAllUsers find the list of the users (only for admin)
func FindAllUsers() ([]model.User, error) {
	var users []model.User
	resp := postgres.DB.Find(&users)
	if resp.Error != nil {
		return nil, shared.InternalErr{}
	}
	return users, nil
}
