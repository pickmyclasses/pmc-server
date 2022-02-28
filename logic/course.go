package logic

import (
	"errors"
	"fmt"
	"strconv"

	classDao "pmc_server/dao/postgres/class"
	dao "pmc_server/dao/postgres/course"
	reviewDao "pmc_server/dao/postgres/review"
	"pmc_server/model"
	"pmc_server/model/dto"
)

func GetCourseList(pn, pSize int) ([]dto.Course, int64, error) {
	courseList, err := dao.GetCourses(pn, pSize)

	if err != nil {
		return nil, -1, fmt.Errorf("unable to get the list of course: %+v\n", err)
	}

	total, err := dao.GetCourseTotal()
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

		courseDto := dto.Course{
			Course:        &course,
			Classes:       *classList,
			OverallRating: rating.OverAllRating,
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
	course, err := dao.GetCourseByID(idInt)
	if err != nil {
		return nil, err
	}

	classList, err := classDao.GetClassByCourseID(course.ID)
	if err != nil {
		return &dto.Course{
			Course:  course,
			Classes: nil,
		}, nil
	}
	return &dto.Course{
		Course:  course,
		Classes: *classList,
	}, nil
}

func GetClassListByCourseID(id string) (*[]model.Class, int64, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, 0, errors.New("provided ID is invalid")
	}
	classList, total := dao.GetClassListByCourseID(idInt)
	return classList, total, nil
}

func GetCoursesBySearch(courseParam model.CourseFilterParams) ([]model.Course, int64, error) {
	courseList, total, err := dao.GetCoursesBySearch(courseParam)
	if err != nil {
		return nil, -1, err
	}
	return courseList, total, err
}
