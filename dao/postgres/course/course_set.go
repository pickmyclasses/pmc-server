package dao

import (
	"errors"
	"fmt"

	"pmc_server/model"
	"pmc_server/shared"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// CourseSetOps defines a set of operations on the CourseSet db.
type CourseSetOps interface {
	InsertCourseSet(name string, isLeaf bool, courseIDList []int64,
		parentSetID int32, linkedToMajor bool, courseRequired int32) error
	QueryCourseSetByID(id int32) (*model.CourseSet, error)
	QueryCourseSetByName(name string) (*model.CourseSet, error)
	QueryCourseSetList() ([]model.CourseSet, error)
	QueryParentCourseSet(id int32) (*model.CourseSet, error)
	QueryChildrenCourseSetList(id int32) ([]model.CourseSet, error)
	QueryCourseListInCourseSetByID(id int32) ([]model.Course, error)
	QueryCourseListInCourseSetByName(name string) ([]model.Course, error)
	QueryDirectMajorCourseSetsByMajorID(id int32) ([]model.CourseSet, error)
	DeleteCourseSetByID(id int32) error
	DetachCourSetFromParentSetByID(id int32) error
	DetachCourseSetChildrenSetByID(id int32) error
}

// CourseSet defines a query on the CourseSet db.
// majorName - the name of major this course set belongs to.
// querier - the query will be used in the CourseSet db.
type CourseSet struct {
	MajorID int32
	Querier *gorm.DB
}

// InsertCourseSet inserts a course set to the CourseSet database.
// name - name of the course set.
// isLeaf - is the course set a leaf node (ie, no sub course set).
// courseIDList - the course id list this course set contains.
func (c CourseSet) InsertCourseSet(name string, isLeaf bool,
	courseIDList []int64, parentSetID int32, linkedToMajor bool, courseRequired int32) error {

	set := model.CourseSet{
		Name:           name,
		IsLeaf:         isLeaf,
		CourseIDList:   courseIDList,
		ParentSetID:    parentSetID,
		MajorID:        c.MajorID,
		LinkedToMajor:  linkedToMajor,
		CourseRequired: courseRequired,
	}

	res := c.Querier.Create(&set)
	if res.Error != nil || res.RowsAffected == 0 {
		return shared.InternalErr{
			Msg: fmt.Sprintf("Failed to create course set with name %s", name),
		}
	}
	return nil
}

// QueryCourseSetByID queries a course set by the given ID in the database.
// id - the id of the course set we are querying.
func (c CourseSet) QueryCourseSetByID(id int32) (*model.CourseSet, error) {
	var set model.CourseSet
	res := c.Querier.Where("id = ?", id).First(&set)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("failed to query course set with the id %d", id),
		}
	}
	return &set, nil
}

// QueryCourseSetByName queries a course set by the given name in the database.
// name - the name of the course set we are querying.
func (c CourseSet) QueryCourseSetByName(name string) (*model.CourseSet, error) {
	var set model.CourseSet
	res := c.Querier.Where("name = ?", name).First(&set)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to query course set with the name %s", name),
		}
	}
	return &set, nil
}

// QueryCourseSetList queries the entire course set list in the database
func (c CourseSet) QueryCourseSetList() ([]model.CourseSet, error) {
	var setList []model.CourseSet
	res := c.Querier.Find(&setList)
	if res.Error != nil {
		return nil, shared.InternalErr{
			Msg: "Failed to fetch the course set list",
		}
	}
	return setList, nil
}

// QueryParentCourseSet queries the parent course set of the set with the given id.
// The purpose of this function is to fetch if the current course set belongs to a specific course set.
// id - the id of current course set we are fetching, not the parent course set id.
func (c CourseSet) QueryParentCourseSet(id int32) (*model.CourseSet, error) {
	var set model.CourseSet
	res := c.Querier.Where("id = ?", id).First(&set)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to query parent course set of the given set with the id %d", id),
		}
	}

	var parentSet model.CourseSet
	res = c.Querier.Where("parent_set_id = ?", set.ParentSetID).First(&parentSet)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to query parent course set of the given set with the id %d", id),
		}
	}
	return &set, nil
}

// QueryChildrenCourseSetList queries the children course list by the given id.
// id - the id of the parent course set.
func (c CourseSet) QueryChildrenCourseSetList(id int32) ([]model.CourseSet, error) {
	var children []model.CourseSet
	res := c.Querier.Where("parent_set_id = ?", id).Find(&children)
	if res.Error != nil {
		// it's ok if the course set doesn't really have any child
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to query the children set of the course set with id %d", id),
		}
	}
	return children, nil
}

// QueryCourseListInCourseSetByID queries course list in the course set of the course set with the given id.
// Course set only stores the id's of the courses, so need some conversion here.
// id - the id of the course set.
func (c CourseSet) QueryCourseListInCourseSetByID(id int32) ([]model.Course, error) {
	var set model.CourseSet
	res := c.Querier.Where("id = ?", id).First(&set)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to query course list of the given set with the id %d", id),
		}
	}

	courseList, err := c.queryCourseListByIDs(set.CourseIDList)
	if err != nil {
		return nil, err
	}
	return courseList, nil
}

// QueryCourseListInCourseSetByName queries course list in the course of the course set with the given id.
// name - the name of the course set.
func (c CourseSet) QueryCourseListInCourseSetByName(name string) ([]model.Course, error) {
	var set model.CourseSet
	res := c.Querier.Where("name = ?", name).First(&set)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to query course list of the given set with the id %s", name),
		}
	}
	courseList, err := c.queryCourseListByIDs(set.CourseIDList)
	if err != nil {
		return nil, err
	}
	return courseList, nil
}

// QueryDirectMajorCourseSetsByMajorID queries the course set list directly associated to the major.
// By direct course set, we mean those course set that has no parent set (eg. general education).
// id - the id of the major we are querying
func (c CourseSet) QueryDirectMajorCourseSetsByMajorID(id int32) ([]model.CourseSet, error) {
	var setList []model.CourseSet
	res := c.Querier.Where("major_id = ? and linked_to_major = ? and parent_set_id = ?", id, true, -1).Find(&setList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []model.CourseSet{}, nil
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to query course set list for the major with id %d", id),
		}
	}
	return setList, nil
}

// DeleteCourseSetByID deletes a course set from the database with the given id.
// id - the id of the course set.
func (c CourseSet) DeleteCourseSetByID(id int32) error {
	res := c.Querier.Where("id = ?", id).Delete(&model.CourseSet{})
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return shared.ContentNotFoundErr{}
		}
		return shared.InternalErr{
			Msg: fmt.Sprintf("Failed to delete course of the given set with the id %d", id),
		}
	}
	return nil
}

// DetachCourSetFromParentSetByID detaches the course set from its parent set.
// This operation only detaches the relation between parent set and child set, not deleting.
// id - the id of the child course set we are detaching
func (c CourseSet) DetachCourSetFromParentSetByID(id int32) error {
	res := c.Querier.Model(&model.CourseSet{}).Where("id = ?", id).Update("parent_set_id", nil)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return shared.ContentNotFoundErr{}
		}
		return shared.InternalErr{
			Msg: fmt.Sprintf("Failed to detach from parent set of the given set with the id %d", id),
		}
	}
	return nil
}

// DetachCourseSetChildrenSetByID detaches all the children set from the parent course set with the given id.
// id - the id of the parent course set we are detaching.
func (c CourseSet) DetachCourseSetChildrenSetByID(id int32) error {
	res := c.Querier.Model(&model.CourseSet{}).Where("parent_set_id = ?", id).
		Update("parent_set_id", nil)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return shared.ContentNotFoundErr{}
		}
		return shared.InternalErr{
			Msg: fmt.Sprintf("Failed to detach the children sets of the given set with the id %d", id),
		}
	}
	return nil
}

// queryCourseListByIDs is a helper function for querying the course entity list of the given id list.
// idList - the id list of the course list
func (c CourseSet) queryCourseListByIDs(idList pq.Int64Array) ([]model.Course, error) {
	var courseList []model.Course
	for _, cid := range idList {
		var course model.Course
		res := c.Querier.Where("id = ?", cid).First(&course)
		if res.Error != nil {
			return nil, shared.InternalErr{
				Msg: fmt.Sprintf("Error when fetching course with course id %d", cid),
			}
		}
		courseList = append(courseList, course)
	}
	return courseList, nil
}
