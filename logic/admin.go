package logic

import (
	"pmc_server/dao/postgres/admin"
	model "pmc_server/model"
)

func GetUserList() ([]model.User, error) {
	return dao.FindAllUsers()
}
