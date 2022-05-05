// Package controller - controller for course entities
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package controller

import (
	"errors"
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// CreateBatchCourseInSetParam defines the parameters for  create a course set with a batch of courses
type CreateBatchCourseInSetParam struct {
	CourseNameList      []string `json:"courseNameList"`      // the list of course names
	SetName             string   `json:"setName"`             // the name of the set
	LinkedToMajor       bool     `json:"linkedToMajor"`       // is the course set directly linked to the major
	CourseRequiredInSet int32    `json:"courseRequiredInSet"` // the amount of courses needed to complete the course set
	IsLeaf              bool     `json:"isLeaf"`              // Is the course set a leaf set, or does it have any subsets
	MajorID             int32    `json:"majorID"`             // the id of the major we are fetching
	ParentSetID         int32    `json:"parentSetID"`         // the id of the parent course set of the current course set
}

// NameList defines a collection of course names
type NameList struct {
	CourseNameList []string `json:"courseNameList"`
}

// Entity is a course entity that only contains the basic info
type Entity struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

// GetCourseListHandler Get the entire course list
// @Summary Use this API to get the list of the classes
// @Description This API is used to get the course list, you should do pagination
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Param pn query int false "Page number of the paginated course list, default is 1" minimum(1)
// @Param psize query int false "Page size(number of course to fetch) of the paginated course list, default is 10" minimum(1) maximum(30)
// @Success 200 {int} total number of the courses
// @Success 200 {array} dto.Course
// @Router /course/list [get]
func GetCourseListHandler(c *gin.Context) {
	pnInt, pSizeInt, err := HandlePagination(c, "10")
	if err != nil {
		_ = c.Error(err)
		return
	}

	courseList, total, err := logic.GetCourseList(pnInt, pSizeInt)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA:  courseList,
		TOTAL: total,
	})
}

// GetCourseByIDHandler Get course and its classes by the given ID
// @Summary Use this API to get the class by the given ID
// @Description This API is used to get the course info along with the classes
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Param id query int true "course id"
// @Success 200 {object} dto.Class
// @Router /course/:id [get]
func GetCourseByIDHandler(c *gin.Context) {
	var courseParam model.CourseParams
	var err error
	if err = c.ShouldBindUri(&courseParam); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	id := c.Query("userID")
	var uid int
	if id != "NaN" {
		uid, err = strconv.Atoi(id)
		if err != nil {
			_ = c.Error(shared.ParamIncompatibleErr{
				Msg: "the given user ID is invalid, should be an integer or NaN",
			})
		}
	} else {
		uid = 0
	}

	courseInfo, err := logic.GetCourseInfo(courseParam.ID, int64(uid))

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: &courseInfo,
	})
}

// GetClassesOfCourseHandler Get the class list of the given course
// @Summary Use this API to get the list of the classes
// @Description This API is used to get the class list, you should do pagination
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Param id query int true "course id"
// @Success 200 {int} total number of the courses
// @Success 200 {array} dto.Class
// @Router /course/:id/class [get]
func GetClassesOfCourseHandler(c *gin.Context) {
	var courseParam model.CourseParams
	if err := c.ShouldBindUri(&courseParam); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	classList, total, err := logic.GetClassListByCourseID(courseParam.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		TOTAL: total,
		DATA:  classList,
	})
}

// GetCoursesBySearchHandler Get the entire course list
// @Summary Use this API to get the list of the classes
// @Description This API is used to get the course list, you should do pagination
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Success 200 {int} total number of the courses
// @Success 200 {array} dto.Course
// @Router /course/list [post]
func GetCoursesBySearchHandler(c *gin.Context) {
	var param model.CourseFilterParams

	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	data, total, err := logic.GetCoursesBySearch(param)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA:  data,
		TOTAL: total,
	})
}

// CreateBatchCourseInSetHandler insert a course set with a batch of courses
// @Summary Use this API to insert a course set along with the courses in the list
// @Description This API is used to create a course set, should be used internally
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Router /course/set [post]
func CreateBatchCourseInSetHandler(c *gin.Context) {
	var param CreateBatchCourseInSetParam
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.CreateCourseSet(param.SetName, param.CourseNameList,
		param.LinkedToMajor, param.CourseRequiredInSet, param.IsLeaf, param.MajorID, param.ParentSetID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		shared.MESSAGE: shared.SUCCESS,
	})
}

// GetCourseByNameHandler get the course entity of the single given ID
// @Summary Use this API to fetch the course entity with the given ID
// @Description This API is used to fetch a course entity
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Router /course [get]
func GetCourseByNameHandler(c *gin.Context) {
	courseName := c.Query("name")
	if courseName == "" {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	id, err := logic.GetCourseByName(courseName)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		shared.DATA: id,
	})
}

// GetCourseIDsByNameListHandler gets multiple course entities with the given course ID list
// @Summary Use this API to fetch a list of course entities with the given course ID list
// @Description This API is used to fetch a list of courses
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Router /course [post]
func GetCourseIDsByNameListHandler(c *gin.Context) {
	var nameList NameList
	if err := c.ShouldBindJSON(&nameList); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	courseList := make([]Entity, 0)
	for _, name := range nameList.CourseNameList {
		id, err := logic.GetCourseByName(name)
		if err != nil {
			if errors.Is(err, shared.ContentNotFoundErr{}) {
				continue
			}
			return
		}
		courseList = append(courseList, Entity{
			Name: name,
			ID:   id,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: courseList,
	})
}
