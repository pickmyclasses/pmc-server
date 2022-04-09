package dao

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"pmc_server/model"
	"pmc_server/shared"
)

// MajorOps defines the set of operations supports to be operated with Major db.
type MajorOps interface {
	InsertMajor(string, bool) error
	InsertEmphasis(string, int32) error
	QueryMajorByID(int32) (model.Major, error)
	QueryEmphasisListByMajorID(int32) ([]model.Major, error)
	QueryEmphasisListByMajorName(string) ([]model.Major, error)
	QueryMajorList() ([]model.Major, error)
	QueryMajorByName(string) (model.Major, error)
	DeleteMajorByID(int32) error
}

// Major defines a query to the Major db.
// collegeID - the college ID this Major is associated with.
// querier - the query will be used to interact with the Major db.
type Major struct {
	CollegeID int32
	Querier   *gorm.DB
}

// InsertMajor inserts a Major entity to the db.
// name - name of the major.
// emphasisRequired - does the major require an emphasis.
func (m Major) InsertMajor(name string, emphasisRequired bool) error {
	major := model.Major{
		CollegeID:        m.CollegeID,
		Name:             name,
		EmphasisRequired: emphasisRequired,
	}
	res := m.Querier.Create(&major)
	if res.Error != nil {
		return shared.InternalErr{
			Msg: "Failed to create major",
		}
	}
	return nil
}

// InsertEmphasis inserts an Emphasis entity to the db.
// Note: Emphasis is a special case of Major, difference is is_emphasis field, and it has
// to be associated with a Major entity.
// name - the name of the emphasis.
// mainMajorID - the ID of the major this emphasis is associated to.
func (m Major) InsertEmphasis(name string, mainMajorID int32) error {
	major := model.Major{
		CollegeID:        m.CollegeID,
		Name:             name,
		EmphasisRequired: false,
		IsEmphasis:       true,
		MainMajorID:      mainMajorID,
	}
	res := m.Querier.Create(&major)
	if res.Error != nil {
		return shared.InternalErr{
			Msg: fmt.Sprintf("Failed to create emphasis %s", name),
		}
	}
	return nil
}

// QueryMajorByID queries a Major entity by given its id.
// id - the ID of the major entity.
func (m Major) QueryMajorByID(id int32) (*model.Major, error) {
	var major model.Major
	res := m.Querier.Where("id = ?", id).First(&major)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: "Failed to query the major",
		}
	}
	return &major, nil
}

// QueryEmphasisListByMajorID queries the emphasis list associated to the given Major with the given major id.
// majorID - the ID of the major entity.
func (m Major) QueryEmphasisListByMajorID(majorID int32) ([]model.Major, error) {
	var emphasisList []model.Major
	res := m.Querier.Where("main_major_id = ?", majorID).Find(&emphasisList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: "Failed to query emphasis list for the major",
		}
	}
	return emphasisList, nil
}

// QueryEmphasisListByMajorName queries the emphasis list associated to the given Major with the major name.
// majorName - the name of the major.
func (m Major) QueryEmphasisListByMajorName(majorName string) ([]model.Major, error) {
	var emphasisList []model.Major
	var mainMajor model.Major
	res := m.Querier.Where("name = ?", majorName).First(&mainMajor)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []model.Major{}, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to fetch the emphasis list of major %s", majorName),
		}
	}

	res = m.Querier.Where("main_major_id = ?", mainMajor.ID).Find(&emphasisList)
	if res.Error != nil {
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Failed to fetch the emphasis list of major %s", majorName),
		}
	}
	return emphasisList, nil
}

// QueryMajorList queries for the entire list of the major in the college.
func (m Major) QueryMajorList() ([]model.Major, error) {
	var majorList []model.Major
	res := m.Querier.Find(&majorList)
	if res.Error != nil {
		return nil, shared.InternalErr{
			Msg: "Failed to fetch the major list",
		}
	}
	return majorList, nil
}

// QueryMajorByName queries a Major entity by the given name.
// If the major is not found, a solvable error will be thrown (404).
// name - the name of the major
func (m Major) QueryMajorByName(name string) (*model.Major, error) {
	var major model.Major
	res := m.Querier.Where("name = ?", name).First(&major)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &model.Major{}, shared.ContentNotFoundErr{}
		}
		return nil, shared.InternalErr{
			Msg: fmt.Sprintf("Did not find the major with given name %s", name),
		}
	}
	return &major, nil
}

// DeleteMajorByID deletes a Major entity from database by the given id.
// id - the id of the major.
func (m Major) DeleteMajorByID(id int32) error {
	var major model.Major
	res := m.Querier.Where("id = ?", id).Delete(&major)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return shared.ContentNotFoundErr{}
		}
		return shared.InternalErr{
			Msg: fmt.Sprintf("Did not delete the major with given ID %d", id),
		}
	}

	if res.RowsAffected == 0 {
		return shared.InternalErr{
			Msg: fmt.Sprintf("Did not delete the major with given ID %d", id),
		}
	}
	return nil
}
