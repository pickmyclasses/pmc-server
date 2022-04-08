package logic

import (
	"pmc_server/dao/aura/major"
	"pmc_server/model/dto"
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
	CollegeID        int      `json:"collegeID"`
	Name             string   `json:"name"`
	DegreeHour       int32    `json:"degreeHour"`
	MinMajorHour     int32    `json:"minMajorHour"`
	EmphasisRequired bool     `json:"emphasisRequired"`
	EmphasisList     []string `json:"emphasisList"`
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

type CourseSet struct {
	SetName      string       `json:"setName"`
	CourseNeeded int32        `json:"courseNeeded"`
	CourseList   []dto.Course `json:"courseList"`
	SubSets      []CourseSet  `json:"subSets"`
}

func GetCourseSetListByMajor(majorName string) ([]CourseSet, error) {
	return nil, nil
}
