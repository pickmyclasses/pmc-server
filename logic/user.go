// Package logic  - logic for users
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package logic

import (
	"errors"
	"sort"
	"strconv"

	collegeDao "pmc_server/dao/postgres/college"
	courseDao "pmc_server/dao/postgres/course"
	historyDao "pmc_server/dao/postgres/history"
	majorDao "pmc_server/dao/postgres/major"
	reviewDao "pmc_server/dao/postgres/review"
	tagDao "pmc_server/dao/postgres/tag"
	"pmc_server/dao/postgres/user"
	"pmc_server/init/postgres"
	"pmc_server/libs/jwt"
	libs "pmc_server/libs/snowflake"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
)

// Recommendation is used for the recommended courses on the page for the users
type Recommendation struct {
	CourseCatalogList []CourseCatalog `json:"courseCatalogList"` // a list of catalogs (with their catalog name and courses)
}

//CourseCatalog represents a catalog of course collections
type CourseCatalog struct {
	DirectCourseSetName string       `json:"directCourseSetName"` // the name of the catalog
	CourseList          []dto.Course `json:"courseList"`          // the courses the catalog contains
}

// Register creates a new user based on the given information
// Users' password should be encrypted
func Register(param *model.RegisterParams) error {
	// check if the user already exist
	exist, err := dao.UserExist(param.Email)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("user already exist")
	}

	// generate snowflake ID
	userID := libs.GenID()

	return dao.InsertUser(&model.User{
		UserID:    userID,
		Email:     param.Email,
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Password:  param.Password,
		CollegeID: int64(param.CollegeID),
	})
}

// Login lets the user access their data by the given information
func Login(param *model.LoginParams) (*dto.User, error) {
	// read if the user exist first
	user, err := dao.ReadUser(&model.User{
		Email:    param.Email,
		Password: param.Password,
	})
	if err != nil {
		return nil, err
	}

	// JWT token for the user to be continually logged in
	token, err := jwt.GenToken(user.UserID, user.FirstName, user.LastName)
	if err != nil {
		return nil, err
	}

	college, err := collegeDao.GetCollegeByID(int32(user.CollegeID))
	if err != nil {
		return nil, err
	}

	return &dto.User{
		ID:          user.ID,
		Token:       token,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Role:        user.Role,
		CollegeID:   int32(user.CollegeID),
		Major:       user.Major,
		Emphasis:    user.Emphasis,
		SchoolYear:  user.AcademicYear,
		CollegeName: college.Name,
		Email:       user.Email,
	}, nil
}

// GetUserHistoryCourseList gets the user taken courses info along with the entities of the courses
func GetUserHistoryCourseList(userID int64) ([]dto.Course, error) {
	historyCourseList, err := dao.GetUserHistoryCourseList(userID)
	if err != nil {
		return nil, err
	}

	courseDtoList := make([]dto.Course, 0)
	for _, c := range historyCourseList {
		course, err := courseDao.GetCourseByID(int(c))
		if err != nil {
			return nil, err
		}
		classList, _, err := courseDao.GetClassListByCourseID(int(c))
		if err != nil {
			return nil, err
		}

		rating, err := reviewDao.GetCourseOverallRating(c)
		if err != nil {
			return nil, err
		}

		tags, err := tagDao.GetTagListByCourseID(c)
		if err != nil {
			return nil, err
		}
		maxCreditF, err := strconv.ParseFloat(course.MaxCredit, 32)
		if err != nil {
			maxCreditF = 0
		}
		minCreditF, err := strconv.ParseFloat(course.MinCredit, 32)
		if err != nil {
			minCreditF = 0
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
			MaxCredit:          maxCreditF,
			MinCredit:          minCreditF,
			Classes:            *classList,
			OverallRating:      rating.OverAllRating,
			Tags:               tags,
		}

		// check if the course is in user's major, if yes, add an extra attachment to it
		if userID != 0 {
			user, err := dao.GetUserByID(userID)
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
				courseDto.DegreeCatalogs = make([][]string, 0)
			}

			courseSetQuery := courseDao.CourseSet{
				MajorID: int32(major.ID),
				Querier: postgres.DB,
			}

			majorSetList, err := courseSetQuery.QueryMajorCourseSets()
			if err != nil {
				return nil, err
			}

			// check the degree catalogs
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
		courseDtoList = append(courseDtoList, courseDto)
	}
	return courseDtoList, nil
}

// AddUserCourseHistory adds course histories to users' records
func AddUserCourseHistory(userID, courseID int64) error {
	err := dao.AddUserHistoryCourse(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}

// RemoveUserCourseHistory removes a course form given users' history record
func RemoveUserCourseHistory(userID, courseID int64) error {
	err := dao.RemoveUserHistoryCourse(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}

// PostUserMajor posts the major of the user
// By declaring users' major, the created user entity will also be returned
func PostUserMajor(userID int64, majorName, emphasis string, schoolYear string) (*dto.User, error) {
	exist, err := dao.UserExistByID(userID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, shared.ContentNotFoundErr{}
	}

	// update the user major and year here
	user, err := dao.UpdateUserMajorAndYear(userID, majorName, emphasis, schoolYear)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Role:       user.Role,
		CollegeID:  int32(user.CollegeID),
		Major:      user.Major,
		Emphasis:   user.Emphasis,
		SchoolYear: user.AcademicYear,
	}, nil
}

type kv struct {
	Cid   int64
	Score float32
}

// RecommendCourses calculates the highest rated courses in each catalog
// pick the top 8 courses in each catalog, and return to the user
func RecommendCourses(userID int64) (*Recommendation, error) {
	user, err := dao.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// skip the courses that are already in user's history
	history, err := historyDao.GetUserCourseHistoryList(userID)
	if err != nil {
		return nil, err
	}

	majorQ := majorDao.Major{
		CollegeID: int32(user.CollegeID),
		Querier:   postgres.DB,
	}

	major, err := majorQ.QueryMajorByName(user.Major)
	if err != nil {
		return nil, err
	}

	courseSetQ := courseDao.CourseSet{
		MajorID: int32(major.ID),
		Querier: postgres.DB,
	}

	directSets, err := courseSetQ.QueryDirectMajorCourseSetsByMajorID()
	if err != nil {
		return nil, err
	}

	// get the id and all the sub-sets information
	// basically fetching all the courses in the degree catalog
	idScoreMapping := make(map[string][]int64)
	for _, set := range directSets {
		directSubSets, _ := courseSetQ.QueryChildrenCourseSetList(int32(set.ID))
		for _, subset := range directSubSets {
			if subset.IsLeaf {
				if original, ok := idScoreMapping[set.Name]; ok {
					for _, id := range subset.CourseIDList {
						original = append(original, id)
					}
					idScoreMapping[set.Name] = original
				} else {
					idList := make([]int64, 0)
					for _, id := range subset.CourseIDList {
						idList = append(idList, id)
					}
					idScoreMapping[set.Name] = idList
				}
			} else {
				thirdLayer, _ := courseSetQ.QueryChildrenCourseSetList(int32(subset.ID))
				for _, t := range thirdLayer {
					if original, ok := idScoreMapping[set.Name]; ok {
						for _, id := range t.CourseIDList {
							original = append(original, id)
						}
						idScoreMapping[set.Name] = original
					} else {
						idList := make([]int64, 0)
						for _, id := range t.CourseIDList {
							idList = append(idList, id)
						}
						idScoreMapping[set.Name] = idList
					}
				}
			}
		}
	}

	// calculate the average rating and give the courses a score, we only need top 8 courses in all the catalogs
	courseCatalogList := make([]CourseCatalog, 0)
	for k, v := range idScoreMapping {
		kvList := make([]kv, 0)
		for _, id := range v {
			for _, h := range history {
				if id == h.CourseID {
					continue
				}
			}

			rating, _ := reviewDao.GetCourseOverallRating(id)
			score := (rating.OverAllRating + 15) / float32(rating.TotalRatingCount+5)
			if rating.TotalRatingCount == 0 {
				score = 0
			}
			kvList = append(kvList, kv{
				Cid:   id,
				Score: score,
			})
		}

		kvList = removeDuplicate(kvList)

		// sort the slice
		sort.Slice(kvList, func(i, j int) bool {
			return kvList[i].Score > kvList[j].Score
		})

		courseDtoList := make([]dto.Course, 0)
		for i := 0; i < 8; i++ {
			courseEntity, err := buildCourseDtoEntity(kvList[i].Cid)
			if err != nil {
				return nil, err
			}
			courseDtoList = append(courseDtoList, *courseEntity)
		}

		courseCatalogList = append(courseCatalogList, CourseCatalog{
			DirectCourseSetName: k,
			CourseList:          courseDtoList,
		})
	}

	return &Recommendation{CourseCatalogList: courseCatalogList}, nil
}

func removeDuplicate(kvList []kv) []kv {
	allKeys := make(map[int64]bool)
	var list []kv
	for _, item := range kvList {
		if _, value := allKeys[item.Cid]; !value {
			allKeys[item.Cid] = true
			list = append(list, item)
		}
	}
	return list
}

// buildCourseDtoEntity is a helper function to build a single course dto entity
func buildCourseDtoEntity(cid int64) (*dto.Course, error) {
	course, err := courseDao.GetCourseByID(int(cid))
	if err != nil {
		return nil, err
	}

	maxCredit, err := strconv.ParseFloat(course.MaxCredit, 32)
	if err != nil {
		return nil, err
	}

	minCredit, err := strconv.ParseFloat(course.MinCredit, 32)
	if err != nil {
		return nil, err
	}

	classList, _, err := courseDao.GetClassListByCourseID(int(cid))
	if err != nil {
		return nil, err
	}

	rating, err := reviewDao.GetCourseOverallRating(cid)
	if err != nil {
		return nil, err
	}

	tags, err := tagDao.GetTagListByCourseID(cid)
	if err != nil {
		return nil, err
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
		Tags:               tags,
	}, nil
}

func GetUserBookmarks(userID int64) ([]dto.Course, error) {
	courseIDList, err := dao.GetUserBookmarks(userID)
	if err != nil {
		return nil, err
	}
	courseDtoList := make([]dto.Course, 0)
	for _, id := range courseIDList {
		entity, err := buildCourseDtoEntity(id)
		if err != nil {
			continue
		}
		courseDtoList = append(courseDtoList, *entity)
	}
	return courseDtoList, nil
}

func PostUserBookmark(userID, courseID int64) error {
	err := dao.AddUserBookmark(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserBookmark(userID, courseID int64) error {
	err := dao.DeleteUserBookmark(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}
