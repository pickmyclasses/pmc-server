// Package controller - controller for admin
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package controller

import (
	"net/http"

	"pmc_server/logic"
	. "pmc_server/shared"

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
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		DATA: result,
	})
}
