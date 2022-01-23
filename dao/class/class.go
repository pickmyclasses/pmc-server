package dao

import (
	"errors"

	"pmc_server/init/postgres"
	model "pmc_server/model"
	"pmc_server/utils"
)

func GetClasses(pn, pSize int) (*[]model.Class, int64) {
	var classes []model.Class
	total := postgres.DB.Find(&classes).RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Find(&classes)

	return &classes, total
}

func GetClassInfoByID(id int) (*model.Class, error) {
	var class model.Class
	result := postgres.DB.Where("id = ?", id).First(&class)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no class info found")
	}
	return &class, nil
}
