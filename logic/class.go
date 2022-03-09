package logic

import (
	"strconv"

	"pmc_server/dao/postgres/class"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetClassList(pn, pSize int) (*[]model.Class, int64) {
	return dao.GetClasses(pn, pSize)
}

func GetClassByID(id string) (*model.Class, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, shared.ParamIncompatibleErr{}
	}
	return dao.GetClassInfoByID(idInt)
}
