package logic

import (
	"pmc_server/dao/aura/major"
)

func CreateMajor(collegeID int, name string, degreeHour, minMajorHour int32, emphasisRequired bool) (string, error) {
	insertion := major.Insertion{Major: major.Entity{
		CollegeID:        collegeID,
		Name:             name,
		DegreeHour:       degreeHour,
		MinMajorHour:     minMajorHour,
		EmphasisRequired: emphasisRequired,
	}}

	entity, err := insertion.InsertMajor()
	if err != nil {
		return "", err
	}
	return entity, nil
}

func CreateEmphasis(collegeID int32, name string, majorName string, totalCredit int32) (string, error) {
	insertion := major.EmphasisInsertion{
		Emphasis: major.Emphasis{
			Name:          name,
			TotalCredit:   totalCredit,
			MainMajorName: majorName,
			CollegeID:     collegeID,
		},
	}

	emphasis, err := insertion.InsertEmphasis()
	if err != nil {
		return "", err
	}
	return emphasis, nil
}

type MajorDto struct {
	Name string `json:"name"`
}

func GetMajorList(collegeID int32) ([]major.Entity, error) {
	reader := major.Read{
		CollegeID: collegeID,
	}
	majorList, err := reader.FindAll()
	if err != nil {
		return nil, err
	}
	return majorList, nil
}
