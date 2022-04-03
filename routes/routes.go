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
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.Use(auth.Cors(), err.JsonErrReporter())
	{
		// for user
		r.POST("/register", controller.RegisterHandler)
		r.POST("/login", controller.LoginHandler)

		// for schedule
		r.POST("/schedule", controller.AddUserScheduleHandler)
		r.GET("/schedule", controller.GetUserScheduleHandler)
		r.PUT("/schedule", controller.DeleteUserScheduleHandler)

		// for course
		r.GET("/course/list", controller.GetCourseListHandler)
		r.GET("/course/:id", controller.GetCourseByIDHandler)
		r.GET("/course/:id/class", controller.GetClassesOfCourseHandler)
		r.GET("/course/:id/professor/list", controller.GetProfessorListByCourseID)

		// for course search
		r.POST("/course/search", controller.GetCoursesBySearchHandler)

		// for review
		r.GET("/course/:id/review", controller.GetCourseReviewListHandler)
		r.GET("/course/:id/review/:review_id", controller.GetCourseReviewByIDHandler)
		r.POST("/course/:id/review", controller.PostCourseReviewHandler)
		r.PUT("/course/:id/review", controller.UpdateCourseReviewHandler)

		// for class
		r.GET("/class/list", controller.GetClassListHandler)
		r.GET("/class/:id", controller.GetClassByIDHandler)

		// for tags
		r.GET("/course/tag", controller.GetTagListHandler)
		r.GET("/course/:id/tag", controller.GetTagByCourseIDHandler)
		r.POST("/course/:id/tag", controller.CreateTagByCourseIDHandler)
		r.PUT("/course/:id/tag", controller.VoteTagHandler)

		// for professors
		r.GET("/professor/list", controller.GetProfessorListHandler)

		// for university
		r.GET("/college/list", controller.GetCollegeList)
		r.GET("/college/:id", controller.GetCollegeByID)
		r.GET("/college/:id/building/list", controller.GetCollegeBuildings)
		r.GET("/college/:id/building", controller.GetCollegeBuildingByID)
		r.GET("/college/:id/semester/list", controller.GetCollegeSemesterList)

		// for major
		r.GET("/college/:id/major/list", controller.GetMajorList)
		r.GET("/college/:id/major/:id", controller.GetMajorByID)
		r.POST("/college/:id/major", controller.CreateMajor)

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
