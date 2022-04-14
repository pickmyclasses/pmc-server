package logic

import (
	"fmt"
	majorDao "pmc_server/dao/postgres/major"
	dao "pmc_server/dao/postgres/user"
	"pmc_server/init/postgres"
	"strconv"

	"pmc_server/dao/aura/course"
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
	courseList, err := courseDao.GetCourseList(pn, pSize)

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

func GetCourseInfo(id string, uid int64) (*dto.Course, error) {
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

	courseDto := &dto.Course{
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
	}

	// check if the course is in user's major, if yes, add an extra attachment to it
	if uid != 0 {
		user, err := dao.GetUserByID(uid)
		if err != nil {
			return nil, err
		}

		majorQuery := majorDao.Major{
			CollegeID: int32(user.CollegeID),
			Querier:   postgres.DB,
		}

		major, err := majorQuery.QueryMajorByName(user.Major)
		if err != nil {
			return nil, err
		}

		if major.Name == "" {
			return courseDto, nil
		}

		courseSetQuery := courseDao.CourseSet{
			MajorID: int32(major.ID),
			Querier: postgres.DB,
		}

		majorSetList, err := courseSetQuery.QueryMajorCourseSets()
		if err != nil {
			return nil, err
		}

		degreeCatalogList := make([][]string, 0)
		for _, set := range majorSetList {
			for _, cid := range set.CourseIDList {
				catalogTuple := make([]string, 0, 2)
				if cid == course.ID {
					if set.ParentSetID != -1 {
						parentSet, err := courseSetQuery.QueryCourseSetByID(set.ParentSetID)
						if err == nil && parentSet.Name != "" {
							catalogTuple = append(catalogTuple, parentSet.Name)
						}
					}
					catalogTuple = append(catalogTuple, set.Name)
					degreeCatalogList = append(degreeCatalogList, catalogTuple)
				}
			}
		}

		courseDto.DegreeCatalogs = degreeCatalogList
	}

	return courseDto, nil
}

func GetClassListByCourseID(id string) (*[]model.Class, int64, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, 0, shared.MalformedIDErr{}
	}
	classList, total := courseDao.GetClassListByCourseID(idInt)
	return classList, total, nil
}

func GetCoursesBySearch(courseParam model.CourseFilterParams) ([]dto.Course, int, error) {
	courseBoolQuery := courseEsDao.NewBoolQuery(courseParam.PageNumber, courseParam.PageSize)

	if courseParam.Keyword != "" {
		keyword := shared.SeparateLettersAndNums(courseParam.Keyword)
		courseBoolQuery.QueryByKeywords(keyword)
	}

	// get the courses that fit the search criteria
	courseSearchIDList, total, err := courseBoolQuery.DoSearch()
	if err != nil {
		return nil, 0, err
	}

	courseDtoList := make([]dto.Course, 0)
	for _, id := range *courseSearchIDList {
		courseByID, err := courseDao.GetCourseByID(int(id))
		if err != nil {
			return nil, -1, shared.InternalErr{}
		}

		classList, err := classDao.GetClassByCourseID(id)
		if err != nil {
			return nil, -1, shared.InternalErr{}
		}

		var maxCredit float64
		var minCredit float64
		if max, err := strconv.ParseFloat(courseByID.MaxCredit, 32); err == nil {
			maxCredit = max
		}
		if min, err := strconv.ParseFloat(courseByID.MinCredit, 32); err == nil {
			minCredit = min
		}

		rating, err := reviewDao.GetCourseOverallRating(id)
		if err != nil {
			return nil, -1, shared.InternalErr{}
		}

		tags, err := tagDao.GetTagListByCourseID(id)
		if err != nil {
			return nil, -1, shared.InternalErr{}
		}

		courseDto := dto.Course{
			CourseID:           id,
			IsHonor:            courseByID.IsHonor,
			FixedCredit:        courseByID.FixedCredit,
			DesignationCatalog: courseByID.DesignationCatalog,
			Description:        courseByID.Description,
			Prerequisites:      courseByID.Prerequisites,
			Title:              courseByID.Title,
			CatalogCourseName:  courseByID.CatalogCourseName,
			Component:          courseByID.Component,
			MaxCredit:          maxCredit,
			MinCredit:          minCredit,
			Classes:            *classList,
			OverallRating:      rating.OverAllRating,
			Tags:               tags,
		}

		// check if the course is in user's major, if yes, add an extra attachment to it
		if courseParam.UserID != 0 {
			user, err := dao.GetUserByID(courseParam.UserID)
			if err != nil {
				return nil, -1, err
			}

			majorQuery := majorDao.Major{
				CollegeID: int32(user.CollegeID),
				Querier:   postgres.DB,
			}

			major, err := majorQuery.QueryMajorByName(user.Major)
			if err != nil {
				return nil, -1, err
			}

			if major.Name == "" {
				courseDto.DegreeCatalogs = make([][]string, 0)
			}

			courseSetQuery := courseDao.CourseSet{
				MajorID: int32(major.ID),
				Querier: postgres.DB,
			}

			majorSetList, err := courseSetQuery.QueryMajorCourseSets()
			if err != nil {
				return nil, -1, err
			}

			degreeCatalogList := make([][]string, 0)
			for _, set := range majorSetList {
				for _, cid := range set.CourseIDList {
					catalogTuple := make([]string, 0, 2)
					if cid == id {
						if set.ParentSetID != -1 {
							parentSet, err := courseSetQuery.QueryCourseSetByID(set.ParentSetID)
							if err == nil && parentSet.Name != "" {
								catalogTuple = append(catalogTuple, parentSet.Name)
							}
						}
						catalogTuple = append(catalogTuple, set.Name)
						degreeCatalogList = append(degreeCatalogList, catalogTuple)
					}
				}
			}
			courseDto.DegreeCatalogs = degreeCatalogList
		}
		courseDtoList = append(courseDtoList, courseDto)
	}

	return courseDtoList, int(total), nil
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

type CourseInfo struct {
	Name    string
	SetName string
}

func InsertCoursesToSet(courseInfoList []string, targetName, setName, relationToTarget string,
	linkedToMajor bool, courseRequiredInSet int32) ([]string, error) {
	// first insert the set
	setInsertion := course.InsertSet{
		Set: course.Set{
			Name:           setName,
			Relation:       relationToTarget,
			TargetName:     targetName,
			CourseRequired: courseRequiredInSet,
			LinkedToMajor:  linkedToMajor,
		},
	}

	_, err := setInsertion.InsertCourseSet()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resList := make([]string, 0)
	// insert courses into set
	for _, courseName := range courseInfoList {
		courseEntity, err := courseDao.GetCourseByCatalogName(courseName)
		// normally this won't cause any error as long as the course exist
		// if the course doesn't exist, we can't really do anything about it
		if err != nil {
			continue
		}
		entityInsertion := course.InsertEntity{
			Entity: course.Entity{
				Name:    courseName,
				ID:      courseEntity.ID,
				SetName: setName,
			},
		}

		entity, err := entityInsertion.Insert()
		if err != nil {
			return nil, err
		}
		resList = append(resList, entity)
	}

	return resList, nil
}
