package schedule

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "pmc_server/consts"
	"pmc_server/logic"
	"pmc_server/model"
)

func AddUserSchedule(c *gin.Context) {
	var param model.PostScheduleParams
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: BAD_PARAM_ERR,
		})
		return
	}

	err := logic.CreateSchedule(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			MESSAGE: INTERNAL_ERR,
			ERROR:   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

func GetUserSchedule(c *gin.Context) {
	var param model.GetScheduleParams
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: BAD_PARAM_ERR,
		})
		return
	}

	data, err := logic.GetSchedule(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			MESSAGE: INTERNAL_ERR,
			ERROR:   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    data,
	})
}

func DeleteUserSchedule(c *gin.Context) {
	var param model.DeleteScheduleParams
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: BAD_PARAM_ERR,
		})
		return
	}

	err := logic.DeleteSchedule(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			MESSAGE: INTERNAL_ERR,
			ERROR:   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
