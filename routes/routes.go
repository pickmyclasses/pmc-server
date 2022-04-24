// Package routes - routes for the APIs
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package routes

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"pmc_server/controller"
	_ "pmc_server/docs"
	"pmc_server/init/logger"
	"pmc_server/middlewares/auth"
	"pmc_server/middlewares/err"
)

func SetUp(mode string) *gin.Engine {
	// check the current running mode
	// TODO: migrate everything to environment variables
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// inject the logger, the recovery, so the server won't just crash
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.Use(auth.Cors(), err.JsonErrReporter())
	{
		// for user
		r.POST("/register", controller.RegisterHandler)
		r.POST("/login", controller.LoginHandler)
		r.POST("/user/history", controller.AddUserHistoryHandler)
		r.GET("/user/:id/history", controller.GetUserHistoryHandler)
		r.PUT("/user/history", controller.RemoveUserHistoryHandler)
		r.POST("/user/major", controller.PostUserMajorHandler)
		r.GET("/user/:id/recommend", controller.GetUserRecommendCourseHandler)

		// for schedule
		r.POST("/schedule", controller.AddUserScheduleHandler)
		r.GET("/schedule", controller.GetUserScheduleHandler)
		r.PUT("/schedule", controller.DeleteUserScheduleHandler)

		// for course
		r.GET("/course/list", controller.GetCourseListHandler)
		r.GET("/course/:id", controller.GetCourseByIDHandler)
		r.GET("/course/:id/class", controller.GetClassesOfCourseHandler)
		r.GET("/course/:id/professor/list", controller.GetProfessorListByCourseIDHandler)
		r.GET("/course", controller.GetCourseByNameHandler)
		r.POST("/course", controller.GetCourseIDsByNameListHandler)

		// for course search
		r.POST("/course/search", controller.GetCoursesBySearchHandler)

		// for review
		r.GET("/course/:id/review", controller.GetCourseReviewListHandler)
		r.GET("/course/:id/review/:review_id", controller.GetCourseReviewByIDHandler)
		r.POST("/course/:id/review", controller.PostCourseReviewHandler)
		r.PUT("/course/:id/review", controller.UpdateCourseReviewHandler)
		r.POST("/course/:id/review/vote", controller.VoteCourseReviewHandler)
		r.GET("/course/:id/review/user", controller.GetUserInfoOfCourseReviewHandler)

		// for class
		r.GET("/class/list", controller.GetClassListHandler)
		r.GET("/class/:id", controller.GetClassByIDHandler)

		// for tags
		r.GET("/course/tag", controller.GetTagListHandler)
		r.GET("/course/:id/tag", controller.GetTagByCourseIDHandler)
		r.POST("/course/:id/tag", controller.CreateTagByCourseIDHandler)
		r.PUT("/course/:id/tag", controller.VoteTagHandler)

		// for course set
		r.POST("/course/set", controller.CreateBatchCourseInSetHandler)

		// for professors
		r.GET("/professor/list", controller.GetProfessorListHandler)

		// for university
		r.GET("/college/list", controller.GetCollegeListHandler)
		r.GET("/college/:id", controller.GetCollegeByIDHandler)
		r.GET("/college/:id/building/list", controller.GetCollegeBuildingsHandler)
		r.GET("/college/:id/building", controller.GetCollegeBuildingByIDHandler)
		r.GET("/college/:id/semester/list", controller.GetCollegeSemesterListHandler)

		// for major
		r.GET("/college/:id/major/list", controller.GetMajorListHandler)
		r.GET("/college/:id/major/:id", controller.GetMajorByIDHandler)
		r.POST("/college/:id/major", controller.CreateMajorHandler)
		r.POST("/college/:id/emphasis", controller.CreateEmphasisHandler)
		r.GET("/college/:id/emphasis", controller.GetEmphasisHandler)
		r.GET("/college/:id/major/course/list", controller.GetMajorCourseSetHandler)
		r.GET("/college/:id/major/course/requirements", controller.GetMajorDirectRequirementsHandler)

		// for stats
		r.GET("/stats/course/:id/professor/rank", controller.GetCourseProfessorRankingHandler)
		r.GET("/stats/course/:id/load", controller.GetCourseAverageLoadHandler)
		r.GET("/stats/course/:id/rating/trend", controller.GetCourseAverageRatingTrendBySemesterHandler)
		//r.GET("/stats/major/:id/grade", controller.GetMajorTopAvgGradeHandler)

		// for testing
		r.GET("/ping", auth.JWT(), func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	// for swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Content not found",
		})
	})
	return r
}
