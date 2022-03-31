package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

func AddUserScheduleHandler(c *gin.Context) {
	scheduleTypeParam := c.Query("type")
	scheduleType, err := strconv.Atoi(scheduleTypeParam)
	if err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	var param model.PostEventParam
	if err := c.Bind(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	switch scheduleType {
	case 1:
		err := logic.CreateSchedule(param)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			MESSAGE: SUCCESS,
		})
	case 2:
		fmt.Println(param)
		err := logic.CreateCustomEvent(param)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			MESSAGE: SUCCESS,
		})
	default:
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}
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
