package logic

import (
	"fmt"
	"strconv"

	classDao "pmc_server/dao/postgres/class"
	courseDao "pmc_server/dao/postgres/course"
	reviewDao "pmc_server/dao/postgres/review"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
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
		return nil, shared.ParamIncompatibleErr{}
	}

	course, err := courseDao.GetCourseByID(idInt)
	if err != nil {
		return nil, err
	}

	classList, err := classDao.GetClassByCourseID(course.ID)
	if err != nil {
		return nil, err
	}

	rating, err := reviewDao.GetCourseOverallRating(course.ID)
	if err != nil {
		return nil, err
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
		return nil, 0, shared.MalformedIDErr{}
	}
	classList, total := courseDao.GetClassListByCourseID(idInt)
	return classList, total, nil
}

func GetCoursesBySearch(courseParam model.CourseFilterParams) ([]model.Class, int64, error) {
	data, err := classDao.GetClassListByOfferDate([]int{2, 3})
	if err != nil {
		return nil, 0, err
	}

	return *data, 0, nil
	//courseBoolQuery := courseEsDao.NewBoolQuery(courseParam.PageNumber, courseParam.PageSize)
	//
	//if courseParam.Keyword != "" {
	//	courseBoolQuery.QueryByKeywords(courseParam.Keyword)
	//}
	//if courseParam.MinCredit != 0 {
	//	courseBoolQuery.QueryByMinCredit(courseParam.MinCredit)
	//}
	//if courseParam.MaxCredit != 0 {
	//	courseBoolQuery.QueryByMaxCredit(courseParam.MaxCredit)
	//}
	//
	//// get the courses that fit the search criteria
	//courseFitIDList, total, err := courseBoolQuery.DoSearch()
	//if err != nil {
	//	return nil, -1, fmt.Errorf("error when fecthing by keywords %+v", err)
	//}
	//
	//classBoolQuery := classEsDao.NewBoolQuery(courseParam.PageSize)
	//
	//if courseParam.OfferedOnline {
	//	classBoolQuery.QueryByIsOnline(true)
	//}
	//if courseParam.OfferedOffline {
	//	classBoolQuery.QueryByIsInPerson(true)
	//}
	//if courseParam.OfferedHybrid {
	//	classBoolQuery.QueryByIsHybrid(true)
	//}
	//if courseParam.OfferedIVC {
	//	classBoolQuery.QueryByIsIVC(true)
	//}
	//
	//if len(courseParam.Weekday) != 0 {
	//	classBoolQuery.QueryByOfferDates(courseParam.Weekday)
	//}
	//if len(courseParam.TaughtProfessor) != 0 {
	//	for _, professor := range courseParam.TaughtProfessor {
	//		classBoolQuery.QueryByProfessor(professor)
	//	}
	//}
	//
	//var startTime float32
	//var endTime float32
	//if courseParam.StartTime != 0 || courseParam.EndTime != 0 {
	//	if courseParam.StartTime == 0 {
	//		startTime = 0
	//	} else {
	//		startTime = courseParam.StartTime
	//	}
	//	if courseParam.EndTime == 0 {
	//		endTime = 24
	//	} else {
	//		endTime = courseParam.EndTime
	//	}
	//	classBoolQuery.QueryByOfferTime(startTime, endTime)
	//}
	//
	//classFitIDList, total, err := classBoolQuery.DoSearch()
	//
	//intersection := shared.Intersection(*courseFitIDList, *classFitIDList)
	//
	////courseList := make([]dto.Course, 0)
	//
	//return intersection, total, nil
}
