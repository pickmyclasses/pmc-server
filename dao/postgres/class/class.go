package dao

import (
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
)

func GetClasses(pn, pSize int) (*[]model.Class, int64) {
	var classes []model.Class
	total := postgres.DB.Find(&classes).RowsAffected
	postgres.DB.Scopes(shared.Paginate(pn, pSize)).Find(&classes)

	return &classes, total
}

func GetClassInfoByID(id int) (*model.Class, error) {
	var class model.Class
	result := postgres.DB.Where("id = ?", id).First(&class)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &class, nil
}

func GetClassByCourseID(courseID int64) (*[]model.Class, error) {
	var classes []model.Class
	result := postgres.DB.Where("course_id = ?", courseID).Find(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

func GetClassByComponent(components []string) (*[]model.Class, error) {
	var classes []model.Class
	sql := "select * from class where component "
	if shared.Contains(components, "In Person") {
		sql += ""
	}

	return &classes, nil
}






