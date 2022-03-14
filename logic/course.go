package logic

import (
	"fmt"
	"strconv"

	courseEsDao"pmc_server/dao/es/course"
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

	// get the courses that fit the search criteria
	courseFitIDList, total, err := courseBoolQuery.DoSearch()

	if err != nil {
		return nil, -1, fmt.Errorf("error when fecthing by keywords %+v", err)
	}

	components := make([]string, 0)
	if courseParam.OfferedOnline {
		components = append(components, "Online")
	}
	if courseParam.OfferedOffline {
		components = append(components, "In Person")
	}
	if courseParam.OfferedIVC {
		components = append(components, "IVC")
	}
	if courseParam.OfferedHybrid {
		components = append(components, "Hybrid")
	}

	courseResultList := make([]model.Class, 0)
	for _, id := range *courseFitIDList {
		classList, err := classDao.GetClassListByComponent(components, id)
		if err != nil {
			return nil, 0, err
		}

		for _, c := range *classList {
			courseResultList = append(courseResultList, c)
		}

		if len(courseParam.Weekday) != 0 {
			newList, err := classDao.GetClassListByOfferDate(courseParam.Weekday, id)
			if err != nil {
				return nil, 0, err
			}

			courseResultList = Intersection(*newList, courseResultList)
		}

		if courseParam.StartTime != 0 && courseParam.EndTime != 0 {
			newList, err := classDao.GetClassListByTimeslot(courseParam.StartTime, courseParam.EndTime, id)
			if err != nil {
				return nil, 0, err
			}
			courseResultList = Intersection(*newList, courseResultList)
		}

		if len(courseParam.TaughtProfessor) != 0 {
			newList, err := classDao.GetClassListByProfessorNames(courseParam.TaughtProfessor, id)
			if err != nil {
				return nil, 0, err
			}
			courseResultList = Intersection(*newList, courseResultList)
		}
	}

	courseResultIDList := make([]int64, 0)
	for _, res := range courseResultList {
		courseResultIDList = append(courseResultIDList, res.CourseID)
	}

	return courseResultIDList, total, nil
}

func Intersection(s1, s2 []model.Class) (courseIDInter []model.Class) {
	hash := make(map[int64]bool)
	for _, e := range s1 {
		hash[e.CourseID] = true
	}
	for _, e := range s2 {
		if hash[e.CourseID] {
			courseIDInter = append(courseIDInter, e)
		}
	}
	courseIDInter = removeDups(courseIDInter)
	return
}

func removeDups(elements []model.Class)(nodups []model.Class) {
	encountered := make(map[int64]bool)
	for _, element := range elements {
		if !encountered[element.CourseID] {
			nodups = append(nodups, element)
			encountered[element.CourseID] = true
		}
	}
	return
}