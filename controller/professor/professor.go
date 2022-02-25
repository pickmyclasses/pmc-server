package professor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "pmc_server/consts"
	"pmc_server/logic"
)

func GetProfessorList(ctx *gin.Context) {
	professorList, err := logic.GetProfessorList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: professorList,
		MESSAGE: SUCCESS,
	})
}
