package dao

import (
	"fmt"
	"sort"

	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"

	"gorm.io/gorm"
)

type Query struct {
	db *gorm.DB
}

func NewQuery(db *gorm.DB) *Query {
	return &Query{
		db: db,
	}
}

func (c *Query) FilterByCourseID(courseID int64) {
	c.db = c.db.Where("course_id = ?", courseID)
}

func (c *Query) FilterByComponent(component string) {
	c.db = c.db.Where("component = %s", component)
}

func (c *Query) FilterByOfferDates(offerDates []int) {
	sort.Slice(offerDates, func(i, j int) bool {
		return offerDates[i] < offerDates[j]
	})

	dates := shared.ConvertSliceToDateString(offerDates)
	c.db = c.db.Where("offer_date = %s ", dates)
}

func (c *Query) FilterByTimeslot(startTime, endTime float32) {
	c.db = c.db.Where("start_time_float >= %f and end_time_float <= %f", startTime, endTime)
}

func (c *Query) Do() ([]model.Class, error) {
	var classList []model.Class
	res := c.db.Find(&classList)
	if res.Error != nil {
		return nil, shared.InternalErr{}
	}
	return classList, nil
}

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

func GetClassListByComponent(components []string, courseID int64) (*[]model.Class, error) {
	var classes []model.Class
	var sql string
	if len(components) == 0 {
		if courseID != 0 {
			result := postgres.DB.Where("course_id = ?", courseID).Find(&classes)
			if result.Error != nil {
				return nil, shared.InternalErr{}
			}
			return &classes, nil
		}
		return &[]model.Class{}, nil
	}

	if courseID != 0 {
		sql = fmt.Sprintf("select * from class where course_id = %d and component = ", courseID)
	} else {
		sql = "select * from class where component = "
	}

	for i, c := range components {
		if i == len(components)-1 {
			sql += fmt.Sprintf("'%s'", c)
		} else {
			sql += fmt.Sprintf("'%s' or component = ", c)
		}
	}

	result := postgres.DB.Raw(sql).Scan(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

func GetClassListByOfferDate(offerDates []int, courseID int64) (*[]model.Class, error) {
	// sort the dates first to convert to the correct format
	sort.Slice(offerDates, func(i, j int) bool {
		return offerDates[i] < offerDates[j]
	})

	dates := shared.ConvertSliceToDateString(offerDates)

	var classes []model.Class
	var sql string
	if courseID != 0 {
		sql = fmt.Sprintf("course_id = %d and offer_date = ?", courseID)
	} else {
		sql = "offer_date = ?"
	}
	result := postgres.DB.Where(sql, dates).Find(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

func GetClassListByTimeslot(startTime, endTime float32, courseID int64) (*[]model.Class, error) {
	var classes []model.Class
	var sql string
	if courseID != 0 {
		sql = fmt.Sprintf("course_id = %d and start_time_float >= ? and end_time_float <= ?", courseID)
	} else {
		sql = "start_time_float >= ? and end_time_float <= ?"
	}

	result := postgres.DB.Where(sql, startTime, endTime).Find(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

func GetClassListByProfessorNames(professorNames []string, courseID int64) (*[]model.Class, error) {
	var professors []int32
	sql := "select id from professor where name = "
	for i, c := range professorNames {
		if i == len(professorNames)-1 {
			sql += fmt.Sprintf("'%s'", c)
		} else {
			sql += fmt.Sprintf("'%s' or name = ", c)
		}
	}
	result := postgres.DB.Raw(sql).Scan(&professors)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}

	if result.RowsAffected == 0 {
		return &[]model.Class{}, nil
	}

	var classes []model.Class
	var classFetchSql string

	if courseID != 0 {
		classFetchSql = fmt.Sprintf("select * from class where course_id = %d and instructor_id = ", courseID)
	} else {
		classFetchSql = "select * from class where instructor_id = "
	}

	for i, p := range professors {
		if i == len(professorNames)-1 {
			classFetchSql += fmt.Sprintf("%d", p)
		} else {
			classFetchSql += fmt.Sprintf("%d or instructor_id = ", p)
		}
	}
	result = postgres.DB.Raw(classFetchSql).Scan(&classes)
	return &classes, nil
}
