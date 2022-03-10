package controller

import (
	"net/http"

	"pmc_server/logic"
	. "pmc_server/shared"

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
		DATA: professorList,
	})
}
