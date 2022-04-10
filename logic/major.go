package logic

import (
	"fmt"
	reviewDao "pmc_server/dao/postgres/review"
	tagDao "pmc_server/dao/postgres/tag"
	"strconv"

	"pmc_server/dao/aura/major"
	classDao "pmc_server/dao/postgres/class"
	dao "pmc_server/dao/postgres/course"
	majorDao "pmc_server/dao/postgres/major"
	"pmc_server/init/postgres"
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

func CreateCourseSet(name string, courseNameList []string, linkedToMajor bool, courseRequired int32,
	isLeaf bool, majorID int32, parentSetID int32) error {
	courseIDList := make([]int64, 0)
	for _, name := range courseNameList {
		c, err := dao.GetCourseByCatalogName(name)
		if err != nil {
			fmt.Printf("course %s does not exist and quit for %+v\n", name, err)
			continue
		}
		courseIDList = append(courseIDList, c.ID)
	}

	courseSetQuery := dao.CourseSet{
		MajorID: majorID,
		Querier: postgres.DB,
	}

	err := courseSetQuery.InsertCourseSet(name, isLeaf, courseIDList, parentSetID, linkedToMajor, courseRequired)
	if err != nil {
		return err
	}
	return nil
}

func GetMajorCourseSets(collegeID int32, majorName string) ([]CourseSet, error) {
	MajorQuery := majorDao.Major{
		CollegeID: collegeID,
		Querier:   postgres.DB,
	}

	m, err := MajorQuery.QueryMajorByName(majorName)
	if err != nil {
		return nil, err
	}

	q := dao.CourseSet{
		MajorID: int32(m.ID),
		Querier: postgres.DB,
	}

	directCourseSetList, err := q.QueryDirectMajorCourseSetsByMajorID(q.MajorID)
	if err != nil {
		return nil, err
	}

	courseSetDtoList := make([]CourseSet, 0)
	for _, set := range directCourseSetList {
		courseEntityList := make([]dto.Course, 0)
		for _, id := range set.CourseIDList {
			course, err := dao.GetCourseByID(int(id))
			if err != nil {
				return nil, err
			}
			classList, err := classDao.GetClassByCourseID(id)
			if err != nil {
				return nil, err
			}

			rating, err := reviewDao.GetCourseOverallRating(id)
			if err != nil {
				return nil, err
			}

			tagList, err := tagDao.GetTagListByCourseID(id)
			if err != nil {
				return nil, err
			}

			maxCredit, err := strconv.Atoi(course.MaxCredit)
			if err != nil {
				maxCredit = 0
			}
			minCredit, err := strconv.Atoi(course.MinCredit)
			if err != nil {
				minCredit = 0
			}
			courseEntityList = append(courseEntityList, dto.Course{
				CourseID:           id,
				IsHonor:            course.IsHonor,
				FixedCredit:        course.FixedCredit,
				DesignationCatalog: course.DesignationCatalog,
				Description:        course.Description,
				Prerequisites:      course.Prerequisites,
				Title:              course.Title,
				CatalogCourseName:  course.CatalogCourseName,
				Component:          course.Component,
				MaxCredit:          float64(maxCredit),
				MinCredit:          float64(minCredit),
				Classes:            *classList,
				OverallRating:      rating.OverAllRating,
				Tags:               tagList,
			})
		}
		courseSetDto := CourseSet{
			SetName:      set.Name,
			CourseNeeded: set.CourseRequired,
			CourseList:   courseEntityList,
			SubSets:      nil,
		}
		courseSetDtoList = append(courseSetDtoList, courseSetDto)
	}

	return courseSetDtoList, nil
}
