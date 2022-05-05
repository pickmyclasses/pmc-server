// Package controller - controller for college entities
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

// GetCollegeListHandler API for getting the list of the college we have in database
// @Summary Use this API to fetch the entire list of colleges
// @Description This API will only be used for registration
// @Tags College
// @Accept application/json
// @Produce application/json
// @Success 200 {string} ok
// @Router /college/list [get]
func GetCollegeListHandler(ctx *gin.Context) {
	collegeList, err := logic.GetCollegeList()
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: collegeList,
	})
}

// GetCollegeByIDHandler API for getting the college entity by its ID
// @Summary Use this API to fetch college entity by collegeID
// @Description This API will be used for fetching a college entity
// @Tags College
// @Accept application/json
// @Produce application/json
// @Success 200 {string} ok
// @Param id query int true "college id"
// @Router /college/:id [get]
func GetCollegeByIDHandler(ctx *gin.Context) {

}

// GetCollegeBuildingsHandler API for getting the list of the building list in a college
// @Summary Use this API to fetch the entire list of buildings in the college
// @Description This API will be used for Google Map api for fetching the building address data
// @Tags College
// @Accept application/json
// @Produce application/json
// @Success 200 {string} ok
// @Param id query int true "college id"
// @Router /college/:id/building/list [get]
func GetCollegeBuildingsHandler(ctx *gin.Context) {

}

// GetCollegeBuildingByIDHandler API for getting the building entity in a college
// @Summary Use this API to fetch a building entity by its ID
// @Description This API will be used for fetching a building details (address, name, ect)
// @Tags College
// @Accept application/json
// @Produce application/json
// @Success 200 {string} ok
// @Param id query int true "college id"
// @Query building_id int true "building id"
// @Router /college/:id/building [get]
func GetCollegeBuildingByIDHandler(ctx *gin.Context) {

}

// GetCollegeSemesterListHandler API for getting the list of the semesters of a college
// @Summary Use this API to fetch the entire list of semesters by college ID
// @Description This API will be used for fetching list of semesters
// @Tags College
// @Accept application/json
// @Produce application/json
// @Success 200 {string} ok
// @Param id query int true "college id"
// @Router /college/:id/semester/list [get]
func GetCollegeSemesterListHandler(ctx *gin.Context) {
	var collegeID int
	var err error
	if collegeID, err = strconv.Atoi(ctx.Param("id")); err != nil || collegeID == 0 {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	semesterList, err := logic.GetCollegeSemesterList(int32(collegeID))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: semesterList,
	})
}
