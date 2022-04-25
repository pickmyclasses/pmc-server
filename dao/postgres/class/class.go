package dao

import (
	"errors"
	"fmt"
	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/shared"
	"sort"

	"gorm.io/gorm"
)

// Query defines an entity for fetching classes
// this query is mostly for the filter function
type Query struct {
	Db          *gorm.DB
	queryString string
}

// FilterByCourseID filters out the classes with the given course ID
func (c *Query) FilterByCourseID(courseID int64) {
	c.queryString += fmt.Sprintf("id = %d and ", courseID)
}

// FilterByComponent filters out the classes with the given component (IVC, In person, ect)
func (c *Query) FilterByComponent(component string) {
	c.queryString += fmt.Sprintf("component = %s and ", component)
}

// FilterByOfferDates filters out the classes with the given dates
func (c *Query) FilterByOfferDates(offerDates []int) {
	sort.Slice(offerDates, func(i, j int) bool {
		return offerDates[i] < offerDates[j]
	})

	dates := shared.ConvertSliceToDateString(offerDates)
	c.queryString += fmt.Sprintf("offer_date = %s and ", dates)
}

// FilterByStartTime filters out the classes with the given start time
func (c *Query) FilterByStartTime(startTime float32) {
	c.queryString += fmt.Sprintf("start_time_float >= %f and ", startTime)
}

func (c *Query) FilterByEndTime(endTime float32) {
	c.queryString += fmt.Sprintf("end_time_float <= %f and ", endTime)
}

func (c *Query) FilterByIsOnline() {
	c.queryString += "offer_date = '' and location = '' and "
}

func (c *Query) FilterByIsOffline() {
	c.queryString += "offer_date != '' and location != '' and "
}

func (c *Query) FilterByCourseIDList(majorIDList []int64) {
	courseIDQuery := "course_id in ("
	for idx, id := range majorIDList {
		if idx == len(majorIDList)-1 {
			courseIDQuery += fmt.Sprintf("%d", id)
			continue
		}
		courseIDQuery += fmt.Sprintf("%d,", id)
	}
	courseIDQuery += ")"
	c.queryString += courseIDQuery
}

// DoSearchEntity starts the query with the given filters to search the class entities (ie, select *)
func (c *Query) DoSearchEntity() ([]model.Class, error) {
	var classList []model.Class
	finalQuery := fmt.Sprintf("select * from class where %s", c.queryString)
	res := c.Db.Raw(finalQuery).Scan(&classList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []model.Class{}, nil
		}
		return nil, shared.InternalErr{
			Msg: "failed to search class entity",
		}
	}
	return classList, nil
}

// DoSearchCourseIDList starts the query with the given filter to search the course_id list (ie, select course_id)
func (c *Query) DoSearchCourseIDList() ([]int64, error) {
	var idList []int64
	finalQuery := fmt.Sprintf("select course_id from class where %s", c.queryString)
	fmt.Printf("final query is %s\n", finalQuery)
	res := c.Db.Raw(finalQuery).Scan(&idList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []int64{}, nil
		}
		return nil, shared.InternalErr{
			Msg: "failed to search class entity",
		}
	}
	return idList, nil
}

// GetClasses gives the entire list of class entities
// this API should only be used with page number and page size for pagination
func GetClasses(pn, pSize int) (*[]model.Class, int64) {
	var classes []model.Class
	total := postgres.DB.Find(&classes).RowsAffected
	postgres.DB.Scopes(shared.Paginate(pn, pSize)).Find(&classes)

	return &classes, total
}

// GetClassByID gives a class list with the given class ID
func GetClassByID(id int) (*model.Class, error) {
	var class model.Class
	result := postgres.DB.Where("id = ?", id).First(&class)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, shared.InternalErr{}
	}
	return &class, nil
}

// GetClassByCourseID gives the entire list of classes associated with the given course ID
func GetClassByCourseID(courseID int64) (*[]model.Class, error) {
	var classes []model.Class
	result := postgres.DB.Where("course_id = ?", courseID).Find(&classes)
	if result.Error != nil {
		return nil, shared.InternalErr{}
	}
	return &classes, nil
}

// GetClassListByComponent gives the class list by the given component and course ID
// component can be an empty list
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

// GetClassListByOfferDate give the class list with the given dates and the course ID
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

// GetClassListByTimeslot gives the class list with the given start/end time and course ID
// this is away from the filter, just for singular use
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

// GetClassListByProfessorNames gives the class list with the given professor names and the course ID
// multiple professor names can be passed in
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
