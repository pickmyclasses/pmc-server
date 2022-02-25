package tag

import (
	"github.com/gin-gonic/gin"
	"pmc_server/model"

	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"
)

func GetTagList(ctx *gin.Context) {
	tagList, err := logic.GetTagList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: tagList,
		MESSAGE: SUCCESS,
	})
}

func GetTagsByCourseID(ctx *gin.Context) {
	var param model.GetTagParams
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			ERROR: NO_ID_ERR,
		})
		return
	}

	tagList, err := logic.GetTagOfCourse(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: tagList,
		MESSAGE: SUCCESS,
	})
}
