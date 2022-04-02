package logic

import (
	"pmc_server/dao/postgres/admin"
	"pmc_server/model"
)

func GetUserList() ([]model.User, error) {
	return dao.FindAllUsers()
}
