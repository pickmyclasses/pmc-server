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
	scheduleType := c.Query("type")

	var param model.PostEventParam
	if err := c.Bind(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	switch scheduleType {
	case "class":
		err := logic.CreateSchedule(param)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			MESSAGE: SUCCESS,
		})
	case "custom-event":
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
	scheduleType := c.Query("type")

	var param model.DeleteScheduleParams
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	switch scheduleType {
	case "class":
		err := logic.DeleteSchedule(param.UserID, param.ClassID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			MESSAGE: SUCCESS,
		})
	case "custom-event":
		err := logic.DeleteCustomEvent(param.EventID)
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
