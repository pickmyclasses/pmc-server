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
	CollegeID        int
	Name             string
	DegreeHour       int32
	MinMajorHour     int32
	EmphasisRequired bool
	EmphasisList     []string
}

func GetMajorList(collegeID int32) ([]MajorDto, error) {
	reader := major.Read{
		CollegeID: collegeID,
	}
	majorList, err := reader.FindAll()
	if err != nil {
		return nil, err
	}

	majorDtoList := make([]MajorDto, 0)
	for _, m := range majorList {
		emphasisReader := major.ReadEmphasis{
			CollegeID: collegeID,
			MajorName: m.Name,
		}
		emphasisList, err := emphasisReader.FindAllEmphasisesOfAMajor()
		if err != nil {
			return nil, err
		}
		emphasisNameList := make([]string, 0)
		for _, e := range emphasisList {
			emphasisNameList = append(emphasisNameList, e.Name)
		}

		majorDto := MajorDto{
			CollegeID:        int(collegeID),
			Name:             m.Name,
			DegreeHour:       m.DegreeHour,
			MinMajorHour:     m.MinMajorHour,
			EmphasisRequired: m.EmphasisRequired,
			EmphasisList:     emphasisNameList,
		}

		majorDtoList = append(majorDtoList, majorDto)
	}

	return majorDtoList, nil
}

func GetMajorEmphasisList(collegeID int32, majorName string) ([]major.Emphasis, error) {
	reader := major.ReadEmphasis{
		CollegeID: collegeID,
		MajorName: majorName,
	}
	emphasisList, err := reader.FindAllEmphasisesOfAMajor()
	if err != nil {
		return nil, err
	}
	return emphasisList, nil
}
