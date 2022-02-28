package logic

import (
	"errors"
	"pmc_server/dao/postgres/class"
	"strconv"

	"pmc_server/model"
)

func GetClassList(pn, pSize int) (*[]model.Class, int64) {
	return dao.GetClasses(pn, pSize)
}

func GetClassByID(id string) (*model.Class, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("provided ID is invalid")
	}
	return dao.GetClassInfoByID(idInt)
}
