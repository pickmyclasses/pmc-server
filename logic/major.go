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
