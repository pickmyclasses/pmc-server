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
// @Description This API is used to get the professor list
// @Tags Professor
// @Accept application/json
// @Produce application/json
// @Success 200 {array} dto.Professor
// @Router college/:id/professor/list [get]
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

// GetProfessorListByCourseIDHandler Get professor list of a course
// @Summary Use this API to get the professor list of a course
// @Description This API is used to get the professor list of a course
// @Tags Professor
// @Accept application/json
// @Produce application/json
// @Success 200 {array} dto.Professor
// @Router college/:id/professor/list?course=? [get]
func GetProfessorListByCourseIDHandler(ctx *gin.Context) {
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
