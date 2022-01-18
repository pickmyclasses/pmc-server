package controller

import (
	"fmt"
	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"
	model "pmc_server/models"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var params model.RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			MESSAGE: INSUFFICIENT_PARAM_ERR,
		})
		return
	}

	if err := logic.Register(&params); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			MESSAGE: fmt.Sprintf("Register failed: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

func LoginHandler(c *gin.Context) {
	var params model.LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			MESSAGE: INSUFFICIENT_PARAM_ERR,
		})
		return
	}

	token, err := logic.Login(&params)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			MESSAGE: fmt.Sprintf("Register failed: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    token,
	})
}
