// Package controller - controller for stats API
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
// Change Log: - 4/10/2022 added stats for average load computation for a course
package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// GetCourseProfessorRankingHandler gets the stats of professor ranking list of a course
// @Summary Use this API to get the stats of professor ranking list of the given course
// @Description This API is getting the stats of the professor ranking list of the given course, ranked by semester
// @Tags Stats
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /stats/course/:id/load [get]
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

// GetCourseAverageLoadHandler gets the course average load for the given course
// @Summary Use this API to get the stats of the class load of the given course, by load, we mean the hw, exams, etc.
// @Description This API is for getting the stats of the class load of the given course.
// @Tags Stats
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /stats/course/:id/rating/trend [get]
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

// GetCourseAverageRatingTrendBySemesterHandler gets the stats of average ranking changing trend of the given course
// @Summary Use this API to get the average ranking changing trend of the given course
// @Description This API is for getting teh average ranking changing trend of the given course
// @Tags Stats
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /stats/course/:id/rating/trend [get]
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

func GetCoursePopularity(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id)
	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	popularity, err := logic.GetCoursePopularity(int64(courseID))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		shared.DATA: popularity,
	})
}

// GetMajorTopAvgGradeHandler gets the stats for top 10 rated courses and major average grade
// @Summary Use this API to get the stats for top-rated courses and major average grade
// @Description This API is for getting the average grade of the major and top 10 rated courses
// @Tags Stats
// @Accept application/json
// @Produce application/json
// @Success 200 {string} OK
// @Router /stats/major/:id/grade [get]
//func GetMajorTopAvgGradeHandler(c *gin.Context) {
//	id := c.Param("id")
//	majorID, err := strconv.Atoi(id)
//	if err != nil {
//		_ = c.Error(shared.ParamIncompatibleErr{
//			Msg: "unable to process the given major id",
//		})
//		return
//	}
//
//	highestRated, err := logic.GetHighestRatedMajorGrade(int32(majorID))
//	if err != nil {
//		_ = c.Error(err)
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		shared.DATA: highestRated,
//	})
//}
