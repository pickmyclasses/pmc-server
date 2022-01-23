package logic

import (
	dao "pmc_server/dao/admin"
	model "pmc_server/model"
)

func GetUserList() ([]model.User, error) {
	return dao.FindAllUsers()
}
