package schedule

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "pmc_server/consts"
	"pmc_server/logic"
	"pmc_server/model"
)

func AddUserScheduleHandler(c *gin.Context) {
	var param model.PostScheduleParams
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: BAD_PARAM_ERR,
		})
		return
	}

	err := logic.CreateSchedule(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

func GetUserScheduleHandler(c *gin.Context) {
	var param model.GetScheduleParams
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: BAD_PARAM_ERR,
		})
		return
	}

	data, err := logic.GetSchedule(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ERROR: INTERNAL_ERR,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: data,
	})
}

func DeleteUserScheduleHandler(c *gin.Context) {
	var param model.DeleteScheduleParams
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: BAD_PARAM_ERR,
		})
		return
	}

	err := logic.DeleteSchedule(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			ERROR: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
