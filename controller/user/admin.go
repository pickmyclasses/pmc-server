package controller

import (
	"fmt"
	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"

	"github.com/gin-gonic/gin"
)

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
