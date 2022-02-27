package tag

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	. "pmc_server/consts"
	"pmc_server/logic"
	"pmc_server/model"
)

func GetTagListHandler(ctx *gin.Context) {
	tagList, err := logic.GetTagList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err.Error(),
		})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			ERROR: NO_ID_ERR,
		})
		return
	}

	tagList, err := logic.GetTagOfCourse(int64(courseIDInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA:    tagList,
	})
}

func CreateTagByCourseIDHandler(ctx *gin.Context) {
	var param model.CreateTagParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			ERROR: INSUFFICIENT_PARAM_ERR,
		})
		return
	}

	err := logic.CreateTagByCourseID(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		MESSAGE:    SUCCESS,
	})
}

func VoteTagHandler(ctx *gin.Context) {

}
