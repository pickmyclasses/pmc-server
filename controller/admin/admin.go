package admin

import (
	"fmt"
	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"

	"github.com/gin-gonic/gin"
)

// GetUserListHandler Admin interface for fetching user list
// @Summary Use this API to fetch the entire user list
// @Description This API will only be used by admin for fetching the user list
// @Tags Admin
// @Accept application/json
// @Produce application/json
// @Success 200 {string} ok
// @Router /admin/user/list [get]
func GetUserListHandler(c *gin.Context) {
	result, err := logic.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			MESSAGE: fmt.Sprintf("unable to query user list %v", err),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    result,
	})
}
