package controller

import (
	"net/http"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

func AddUserScheduleHandler(c *gin.Context) {
	var param model.PostScheduleParams
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.CreateSchedule(param)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

func GetUserScheduleHandler(c *gin.Context) {
	var param model.GetScheduleParams
	if err := c.ShouldBind(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	data, err := logic.GetSchedule(param)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: data,
	})
}

func DeleteUserScheduleHandler(c *gin.Context) {
	var param model.DeleteScheduleParams
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.DeleteSchedule(param)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
