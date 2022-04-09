package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/shared"

	"github.com/gin-gonic/gin"
)

func GetCourseProfessorRankingHandler(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id)
	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	professorList, err := logic.GetProfessorRankingByCourseID(int64(courseID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		shared.DATA: professorList,
	})
}

func GetCourseAverageLoadHandler(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id)
	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	courseLoad, err := logic.GetCourseAverageLoad(int64(courseID))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		shared.DATA: courseLoad,
	})
}

func GetCourseAverageRatingTrendBySemesterHandler(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id)
	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	ratingTrend, err := logic.GetCourseRatingTrendBySemester(int64(courseID))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		shared.DATA: ratingTrend,
	})
}
