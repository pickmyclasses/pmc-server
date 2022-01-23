package controller

import (
	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"
	model "pmc_server/model"
	"pmc_server/utils"

	"github.com/gin-gonic/gin"
)

func GetCourseListHandler(c *gin.Context) {
	pnInt, pSizeInt, err := utils.HandlePagination(c, "10")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: err,
		})
	}

	courseList, total := logic.GetCourseList(pnInt, pSizeInt)
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    courseList,
		TOTAL:   total,
	})
}

func GetCourseHandler(c *gin.Context) {
	var courseParam model.CourseParams
	if err := c.ShouldBindUri(&courseParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: NO_ID_ERR,
		})
		return
	}

	courseInfo, err := logic.GetCourseInfo(courseParam.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			MESSAGE: NO_INFO_ERR,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    &courseInfo,
	})
}

func GetClassesOfCourseHandler(c *gin.Context) {
	var courseParam model.CourseParams
	if err := c.ShouldBindUri(&courseParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: NO_ID_ERR,
		})
		return
	}

	pnInt, pSizeInt, err := utils.HandlePagination(c, "20")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: err,
		})
		return
	}

	classList, total, err := logic.GetClassListByCourseID(courseParam.ID, pnInt, pSizeInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			MESSAGE: NO_INFO_ERR,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		TOTAL:   total,
		DATA:    classList,
	})
}
