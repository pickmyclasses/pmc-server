// Package controller - controller for stats Tags
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// GetTagListHandler gets the entire tag list stored in the database
// @Summary Use this API to get the entire tag list stored in the database
// @Description This API is for getting the entire tag list, for every single courses in the database
// @Tags Tags
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /course/tag [get]
func GetTagListHandler(ctx *gin.Context) {
	tagList, err := logic.GetTagList()
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: tagList,
	})
}

// GetTagByCourseIDHandler gets the tag list by the given course ID
// @Summary Use this API to get the tag list by the given course ID, pros and cons
// @Description This API is getting the tag list by the given course ID
// @Tags Tags
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /course/:id/tag [get]
func GetTagByCourseIDHandler(ctx *gin.Context) {
	courseID := ctx.Param("id")
	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	tagList, err := logic.GetTagOfCourse(int64(courseIDInt))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: tagList,
	})
}

// CreateTagByCourseIDHandler creates a tag by the given info for the given course
// @Summary Use this API to post a tag to the given course
// @Description This API is for posting a tag to the given course
// @Tags Tags
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /course/:id/review [post]
func CreateTagByCourseIDHandler(ctx *gin.Context) {
	var param model.CreateTagParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}
	courseID := ctx.Param("id")
	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	err = logic.CreateTagByCourseID(param.Content, param.Type, int64(courseIDInt))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

// VoteTagHandler basically gives the function for the user to upvote tags
// @Summary Use this API to let the users to upvote the tags
// @Description This API is for letting the users to upvote a tag
// @Tags Tags
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /course/:id/tag [get]
func VoteTagHandler(ctx *gin.Context) {
	var param model.VoteTagParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		_ = ctx.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.VoteTag(param.TagID, param.UserID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
