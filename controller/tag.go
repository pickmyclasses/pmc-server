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
