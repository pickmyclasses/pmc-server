package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pmc_server/logic"
	"pmc_server/shared"
	"strconv"
)

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

func GetMajorByIDHandler(ctx *gin.Context) {

}

type CreateParams struct {
	Name             string `json:"name"`
	DegreeHour       int32  `json:"degreeHour" `
	MinMajorHour     int32  `json:"minMajorHour"`
	EmphasisRequired bool   `json:"emphasisRequired"`
}

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

type CreateEmphasisParam struct {
	Name          string `json:"name"`
	TotalCredit   int32  `json:"totalCredit"`
	MainMajorName string `json:"mainMajorName"`
}

func CreateEmphasisHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	collegeID, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamInsufficientErr{})
		return
	}

	var param CreateEmphasisParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(shared.ParamIncompatibleErr{})
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
