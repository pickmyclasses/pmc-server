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
	Name             string `json:"name"`
	EmphasisRequired bool   `json:"emphasisRequired"`
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
		dto := MajorDto{
			Name:             m.Name,
			EmphasisRequired: m.EmphasisRequired,
		}
		majorDtoList = append(majorDtoList, dto)
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
	reader := major.Reader{
		MajorName: majorName,
	}
	courseSetReader := major.ReadList{
		Reader: reader,
	}

	_, err := courseSetReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type DirectCourseSet struct {
	SetName        string `json:"setName"`
	CourseRequired int32  `json:"courseRequired"`
}

func GetDirectMajorCourseSets(majorName string) ([]DirectCourseSet, error) {
	reader := major.Reader{
		MajorName: majorName,
	}
	courseSetReader := major.ReadList{
		Reader: reader,
	}

	_, err := courseSetReader.ReadDirectCourseSet()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetSubCourseSets(majorName string) ([]major.SubSet, error) {
	reader := major.ReadList{
		Reader: major.Reader{
			MajorName: majorName,
			SetName:   "General Education Courses",
		}}
	subsetList, err := reader.ReadSubCourseSets()
	if err != nil {
		return nil, err
	}
	return subsetList, nil
}
