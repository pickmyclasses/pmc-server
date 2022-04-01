package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// GetProfessorListHandler Get the entire professor list
// @Summary Use this API to get the list of the professors
// @Description This API is used to get the course list, you should do pagination
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Success 200 {array} dto.Course
// @Router /course/list [get]
func GetProfessorListHandler(ctx *gin.Context) {
	professorList, err := logic.GetProfessorList()
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: professorList,
	})
}

func GetProfessorListByCourseID(ctx *gin.Context) {
	var courseID int
	var err error
	if courseID, err = strconv.Atoi(ctx.Param("id")); err != nil || courseID == 0 {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}
	professorList, err := logic.GetProfessorListByCourseID(int64(courseID))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: professorList,
	})
}
