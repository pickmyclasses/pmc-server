package logic

import (
	"fmt"
	"strconv"

	courseEsDao "pmc_server/dao/es/course"
	classDao "pmc_server/dao/postgres/class"
	courseDao "pmc_server/dao/postgres/course"
	reviewDao "pmc_server/dao/postgres/review"
	tagDao "pmc_server/dao/postgres/tag"
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

	tagList, err := tagDao.GetTagListByCourseID(course.ID)

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
		Tags:               tagList,
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
		return nil, 0, err
	}
	return *courseFitIDList, total, nil

	//if err != nil {
	//	return nil, -1, fmt.Errorf("error when fetching by keywords %+v", err)
	//}
	//
	//// only check for the filter parameters when there was actually a filter on
	//// this is for saving time from checking the parameters
	//if courseParam.HasFilter {
	//	classQuery := classDao.NewQuery(postgres.DB)
	//
	//	if len(*courseFitIDList) != 0 {
	//		for _, id := range *courseFitIDList {
	//			classQuery.FilterByCourseID(id)
	//		}
	//	}
	//
	//	components := make([]string, 0)
	//	if courseParam.OfferedOnline {
	//		components = append(components, "Online")
	//	}
	//	if courseParam.OfferedOffline {
	//		components = append(components, "In Person")
	//	}
	//	if courseParam.OfferedIVC {
	//		components = append(components, "IVC")
	//	}
	//	if courseParam.OfferedHybrid {
	//		components = append(components, "Hybrid")
	//	}
	//
	//	for _, component := range components {
	//		classQuery.FilterByComponent(component)
	//	}
	//
	//	if courseParam.StartTime != 0 && courseParam.EndTime != 0 {
	//		classQuery.FilterByTimeslot(courseParam.StartTime, courseParam.EndTime)
	//	}
	//
	//	if len(courseParam.Weekday) != 0 {
	//		classQuery.FilterByOfferDates(courseParam.Weekday)
	//	}
	//
	//	classList, err := classQuery.Do()
	//	if err != nil {
	//		return nil, 0, err
	//	}
	//
	//	classFitCourseIDList := make([]int64, 0)
	//	for _, class := range classList {
	//		classFitCourseIDList = append(classFitCourseIDList, class.CourseID)
	//	}
	//
	//	// No search keyword input
	//	if courseParam.Keyword == "" {
	//		courseDtoList, err := buildCourseDto(classFitCourseIDList)
	//		if err != nil {
	//			return nil, 0, err
	//		}
	//		return courseDtoList, total, nil
	//	} else {
	//		finalCourseIDList := intersection(classFitCourseIDList, *courseFitIDList)
	//		courseDtoList, err := buildCourseDto(finalCourseIDList)
	//		if err != nil {
	//			return nil, 0, err
	//		}
	//
	//		return courseDtoList, total, nil
	//	}
	//}
	//
	//courseDtoList, err := buildCourseDto(*courseFitIDList)
	//if err != nil {
	//	return nil, 0, err
	//}
	//
	//return courseDtoList, total, nil
}

func buildCourseDto(idList []int64) ([]dto.Course, error) {
	courseDtoList := make([]dto.Course, 0)
	for _, id := range idList {
		course, err := courseDao.GetCourseByID(int(id))
		if err != nil {
			return nil, shared.InternalErr{}
		}

		classList, err := classDao.GetClassByCourseID(id)
		if err != nil {
			return nil, shared.InternalErr{}
		}

		var maxCredit float64
		var minCredit float64
		if max, err := strconv.ParseFloat(course.MaxCredit, 32); err == nil {
			maxCredit = max
		}
		if min, err := strconv.ParseFloat(course.MinCredit, 32); err == nil {
			minCredit = min
		}

		rating, err := reviewDao.GetCourseOverallRating(id)
		if err != nil {
			return nil, shared.InternalErr{}
		}

		tags, err := tagDao.GetTagListByCourseID(id)
		if err != nil {
			return nil, shared.InternalErr{}
		}

		courseDto := dto.Course{
			CourseID:           id,
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
			Tags:               tags,
		}
		courseDtoList = append(courseDtoList, courseDto)
	}

	return courseDtoList, nil
}

func intersection(s1, s2 []int64) (inter []int64) {
	hash := make(map[int64]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = removeDups(inter)
	return
}

func removeDups(elements []int64) (nodups []int64) {
	encountered := make(map[int64]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}
