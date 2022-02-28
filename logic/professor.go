package logic

import (
	dao "pmc_server/dao/postgres/professor"
	"pmc_server/model"
)

func GetProfessorList() ([]model.Professor, error) {
	return dao.GetProfessorList()
}
