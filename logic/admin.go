package logic

import (
	dao "pmc_server/dao/admin"
	model "pmc_server/models"
)

func GetUserList() ([]model.User, error) {
	return dao.FindAllUsers()
}
