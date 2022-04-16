// Package controller - controller for schedule entity
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package controller

import (
	"net/http"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// AddUserScheduleHandler Post a single schedule for a user
// @Summary Use this API to post a single schedule for a user, (custom event or class event)
// @Description This API is for posting a schedule for a user
// @Tags Schedule
// @Accept application/json
// @Produce application/json
// @Param object body model.PostEventParam true "Post schedule parameters"
// @Success 200 {string} OK
// @Router /course/schedule [post]
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

// GetUserScheduleHandler gets the schedule list of a user
// @Summary Use this API to get the schedule for a user, (custom event or class event)
// @Description This API is for getting the schedule of a user
// @Tags Schedule
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /course/schedule [get]
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

// DeleteUserScheduleHandler removes a single schedule for a user
// @Summary Use this API to remove a single schedule for a user, (custom event or class event)
// @Description This API is for removing a schedule for a user
// @Tags Schedule
// @Accept application/json
// @Produce application/json
// @Param object body model.DeleteScheduleParams true "delete schedule parameters"
// @Success 200 {string} OK
// @Router /course/schedule [put]
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
