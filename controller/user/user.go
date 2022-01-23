package controller

import (
	"fmt"
	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"
	"pmc_server/model"

	"github.com/gin-gonic/gin"
)

// RegisterHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
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

// LoginHandler User login interface
// @Summary Use this API to login
// @Description After login, a token will be returned to verify user in the future
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.LoginParams true "login parameters"
// @Success 200 {string} jwt token
// @Router /login [post]
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
