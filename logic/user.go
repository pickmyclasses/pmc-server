package logic

import (
	"errors"
	courseDao "pmc_server/dao/postgres/course"
	majorDao "pmc_server/dao/postgres/major"
	reviewDao "pmc_server/dao/postgres/review"
	tagDao "pmc_server/dao/postgres/tag"
	"pmc_server/dao/postgres/user"
	"pmc_server/init/postgres"
	"pmc_server/model/dto"
	"pmc_server/shared"
	"strconv"

	"pmc_server/libs/jwt"
	libs "pmc_server/libs/snowflake"
	model "pmc_server/model"
)

func Register(param *model.RegisterParams) error {
	exist, err := dao.UserExist(param.Email)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("user already exist")
	}

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

func Login(param *model.LoginParams) (*dto.User, error) {
	user, err := dao.ReadUser(&model.User{
		Email:    param.Email,
		Password: param.Password,
	})
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenToken(user.UserID, user.FirstName, user.LastName)
	if err != nil {
		return nil, err
	}
	return &dto.User{
		ID:         user.ID,
		Token:      token,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Role:       user.Role,
		CollegeID:  int32(user.CollegeID),
		Major:      user.Major,
		Emphasis:   user.Emphasis,
		SchoolYear: user.AcademicYear,
	}, nil
}

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
		classList, _ := courseDao.GetClassListByCourseID(int(c))
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

func AddUserCourseHistory(userID, courseID int64) error {
	err := dao.AddUserHistoryCourse(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}

func RemoveUserCourseHistory(userID, courseID int64) error {
	err := dao.RemoveUserHistoryCourse(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}

func PostUserMajor(userID int64, majorName, emphasis string, schoolYear string) error {
	exist, err := dao.UserExistByID(userID)
	if err != nil {
		return err
	}
	if !exist {
		return shared.ContentNotFoundErr{}
	}

	err = dao.UpdateUserMajorAndYear(userID, majorName, emphasis, schoolYear)
	if err != nil {
		return err
	}
	return nil
}
