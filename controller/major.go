// Package controller - controller for major entity
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/shared"
	
	"github.com/gin-gonic/gin"
)

// CreateParams defines the action to create a major
type CreateParams struct {
	Name             string `json:"name"`             // the name of the major
	DegreeHour       int32  `json:"degreeHour" `      // the hours the degree required to graduate
	MinMajorHour     int32  `json:"minMajorHour"`     // the minimum major hours required to graduate
	EmphasisRequired bool   `json:"emphasisRequired"` // if the major requires an emphasis
}

// CreateEmphasisParam defines teh action to create an emphasis (also a major)
type CreateEmphasisParam struct {
	Name          string `json:"name"`          // the name of the emphasis
	TotalCredit   int32  `json:"totalCredit"`   // the total credit of the emphasis
	MainMajorName string `json:"mainMajorName"` // the main major name (this emphasis is based in)
}

// GetMajorListHandler gets the entire list of the major in the college
// @Summary Use this API to get the entire list of majors in the college
// @Description This API is used to get the course list, you should do pagination
// @Tags Major
// @Accept application/json
// @Produce application/json
// @Router college/:id/major/list [get]
func GetMajorListHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	collegeID, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamInsufficientErr{})
		return
	}

	majorList, err := logic.GetMajorList(int32(collegeID))
	if err != nil {
		_ = ctx.Error(shared.InternalErr{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: majorList,
	})
}

// GetMajorByIDHandler gets a major entity in the college
// @Summary Use this API to get a major entity by its ID
// @Description This API is used to get a course entity, its ID should be provided
// @Tags Major
// @Accept application/json
// @Produce application/json
// @Router college/:id/major?id= [get]
func GetMajorByIDHandler(ctx *gin.Context) {

}

// CreateMajorHandler create a major entity in the college
// @Summary Use this API to get a major entity, this should be only for internal use
// @Description This API is used to create a course entity
// @Tags Major
// @Accept application/json
// @Produce application/json
// @Router college/:id/major [post]
func CreateMajorHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	collegeID, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamInsufficientErr{})
		return
	}
	var param CreateParams
	if err := ctx.ShouldBindJSON(&param); err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	major, err := logic.CreateMajor(collegeID, param.Name, param.DegreeHour, param.MinMajorHour, param.EmphasisRequired)
	if err != nil {
		_ = ctx.Error(shared.InternalErr{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: major,
	})
}

// CreateEmphasisHandler creates an emphasis entity in the college
// @Summary Use this API to create an emphasis, this is only for internal use
// @Description This API is used to create an emphasis entity
// @Tags Major
// @Accept application/json
// @Produce application/json
// @Router college/:id/emphasis [post]
func CreateEmphasisHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	collegeID, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	var param CreateEmphasisParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	emphasis, err := logic.CreateEmphasis(int32(collegeID), param.Name, param.MainMajorName, param.TotalCredit)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: emphasis,
	})
}

// GetEmphasisHandler gets an emphasis entity in the college
// @Summary Use this API to get an emphasis entity in the college, (emphasis is also a major)
// @Description This API is used to get an emphasis entity. it's also a major, but with a main major
// @Tags Major
// @Accept application/json
// @Produce application/json
// @Router college/:id/emphasis?id= [get]
func GetEmphasisHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	collegeID, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	major := ctx.Query("major")
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	emphasisList, err := logic.GetMajorEmphasisList(int32(collegeID), major)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: emphasisList,
	})
}

// GetMajorCourseSetHandler gets a course set in a major in the college
// @Summary Use this API to get the entire major course set requirements
// @Description This API is used to get the entire major course set requirements (basically everything needed to graduate)
// @Tags Major
// @Accept application/json
// @Produce application/json
// @Router college/:id/major/set?major=? [get]
func GetMajorCourseSetHandler(ctx *gin.Context) {
	major := ctx.Query("major")

	if major == "" {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	courseSetList, err := logic.GetCourseSetListByMajor(major)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: courseSetList,
	})
}

// GetMajorDirectRequirementsHandler gets the direct course set in a major in the college
// @Summary Use this API to get the direct course sets in the major (direct course set means something like gen-eds)
// @Description This API is used to get the direct course set in the major
// @Tags Major
// @Accept application/json
// @Produce application/json
// @Router college/:id/major/requirement [get]
func GetMajorDirectRequirementsHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	collegeID, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	major := ctx.Query("major")
	if major == "" {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	courseSetList, err := logic.GetMajorCourseSets(int32(collegeID), major)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: courseSetList,
	})
}
