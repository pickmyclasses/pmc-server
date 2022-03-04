package logic

import (
	"errors"
	"fmt"
	"strconv"

	courseEsDao "pmc_server/dao/es/course"
	classDao "pmc_server/dao/postgres/class"
	courseDao "pmc_server/dao/postgres/course"
	reviewDao "pmc_server/dao/postgres/review"
	"pmc_server/model"
	"pmc_server/model/dto"
)

func GetCourseList(pn, pSize int) ([]dto.Course, int64, error) {
	courseList, err := courseDao.GetCourses(pn, pSize)

	if err != nil {
		return nil, -1, fmt.Errorf("unable to get the list of course: %+v\n", err)
	}

	total, err := courseDao.GetCourseTotal()
	if err != nil {
		return nil, -1, fmt.Errorf("unable to get the total of course: %+v\n", err)
	}

	courseDtoList := make([]dto.Course, 0)
	for _, course := range courseList {
		classList, err := classDao.GetClassByCourseID(course.ID)
		if err != nil {
			return nil,
				-1,
				fmt.Errorf("failed to fetch class list of the course %s becuase %+v\n", course.CatalogCourseName, err)
		}

		rating, err := reviewDao.GetCourseOverallRating(course.ID)
		if err != nil {
			return nil,
				-1,
				fmt.Errorf("failed to fetch course overall rating of the corse %s becuase %+v\n", course.CatalogCourseName, err)
		}

		maxCredit, err := strconv.ParseFloat(course.MaxCredit, 32)
		if err != nil {
			maxCredit = 0.0
		}
		minCredit, err := strconv.ParseFloat(course.MinCredit, 32)
		if err != nil {
			minCredit = 0.0
		}

		courseDto := dto.Course{
			CourseID:           course.ID,
			IsHonor:            course.IsHonor,
			FixedCredit:        course.FixedCredit,
			DesignationCatalog: course.DesignationCatalog,
			Description:        course.Description,
			Prerequisites:      course.Prerequisites,
			Title:              course.Title,
			CatalogCourseName:  course.CatalogCourseName,
			Component:          course.Component,
			MaxCredit:          maxCredit,
			MinCredit:          minCredit,
			Classes:            *classList,
			OverallRating:      rating.OverAllRating,
		}

		courseDtoList = append(courseDtoList, courseDto)
	}

	return courseDtoList, total, nil
}

func GetCourseInfo(id string) (*dto.Course, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("provided ID is invalid")
	}

	course, err := courseDao.GetCourseByID(idInt)
	if err != nil {
		return nil, err
	}

	classList, err := classDao.GetClassByCourseID(course.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch classes of the course %s: %+v\n", course.CatalogCourseName, err)
	}

	rating, err := reviewDao.GetCourseOverallRating(course.ID)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch the rating of the course %s: %+v\n", course.CatalogCourseName, err)
	}

	maxCredit, err := strconv.ParseFloat(course.MaxCredit, 32)
	if err != nil {
		maxCredit = 0.0
	}
	minCredit, err := strconv.ParseFloat(course.MinCredit, 32)
	if err != nil {
		minCredit = 0.0
	}

	return &dto.Course{
		CourseID:           course.ID,
		IsHonor:            course.IsHonor,
		FixedCredit:        course.FixedCredit,
		DesignationCatalog: course.DesignationCatalog,
		Description:        course.Description,
		Prerequisites:      course.Prerequisites,
		Title:              course.Title,
		CatalogCourseName:  course.CatalogCourseName,
		Component:          course.Component,
		MaxCredit:          maxCredit,
		MinCredit:          minCredit,
		Classes:            *classList,
		OverallRating:      rating.OverAllRating,
	}, nil
}

func GetClassListByCourseID(id string) (*[]model.Class, int64, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, 0, errors.New("provided ID is invalid")
	}
	classList, total := courseDao.GetClassListByCourseID(idInt)
	return classList, total, nil
}

func GetCoursesBySearch(courseParam model.CourseFilterParams) ([]int64, int64, error) {
	courseBoolQuery := courseEsDao.NewBoolQuery(courseParam.PageNumber, courseParam.PageSize)

	if courseParam.Keyword != "" {
		courseBoolQuery.QueryByKeywords(courseParam.Keyword)
	}

	if courseParam.MinCredit != 0 {
		courseBoolQuery.QueryByMinCredit(courseParam.MinCredit)
	}

	if courseParam.MaxCredit != 0 {
		courseBoolQuery.QueryByMaxCredit(courseParam.MaxCredit)
	}

	res, total, err := courseBoolQuery.DoSearch()
	if err != nil {
		return nil, -1, fmt.Errorf("error when fecthing by keywords %+v", err)
	}


	return res, total, nil
}
