package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pmc_server/logic"
	"pmc_server/shared"
	"strconv"
)

func GetMajorList(ctx *gin.Context) {

}

func GetMajorByID(ctx *gin.Context) {

}

type CreateParams struct {
	Name             string `json:"name"`
	DegreeHour       int32  `json:"degreeHour" `
	MinMajorHour     int32  `json:"minMajorHour"`
	EmphasisRequired bool   `json:"emphasisRequired"`
}

func CreateMajor(ctx *gin.Context) {
	id := ctx.Param("id")
	collegeID, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamInsufficientErr{})
		return
	}
	var param CreateParams
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			shared.ERROR: err,
		})
		return
	}

	major, err := logic.CreateMajor(collegeID, param.Name, param.DegreeHour, param.MinMajorHour, param.EmphasisRequired)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			shared.ERROR: err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: major,
	})
}
