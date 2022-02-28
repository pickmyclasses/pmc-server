package routes

import (
	"net/http"

	classController "pmc_server/controller/class"
	courseController "pmc_server/controller/course"
	ProfessorController "pmc_server/controller/professor"
	reviewController "pmc_server/controller/review"
	scheduleController "pmc_server/controller/schedule"
	tagController "pmc_server/controller/tag"
	userController "pmc_server/controller/user"
	_ "pmc_server/docs"
	"pmc_server/init/logger"
	"pmc_server/middlewares/auth"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.Use(auth.Cors())
	{
		// for user
		r.POST("/register", userController.RegisterHandler)
		r.POST("/login", userController.LoginHandler)

		// for schedule
		r.POST("/schedule", scheduleController.AddUserScheduleHandler)
		r.GET("/schedule", scheduleController.GetUserScheduleHandler)
		r.PUT("/schedule", scheduleController.DeleteUserScheduleHandler)

		// for course
		r.GET("/course/list", courseController.GetCourseListHandler)
		r.GET("/course/:id", courseController.GetCourseByIDHandler)
		r.GET("/course/:id/class", courseController.GetClassesOfCourseHandler)

		// for course search
		r.POST("/course/search", courseController.GetCoursesBySearchHandler)

		// for review
		r.GET("/course/:id/review", reviewController.GetCourseReviewListHandler)
		r.GET("/course/:id/review/:review_id", reviewController.GetCourseReviewByIDHandler)
		r.POST("/course/:id/review", reviewController.PostCourseReviewHandler)
		r.PUT("/course/review", reviewController.UpdateCourseReviewHandler)

		// for class
		r.GET("/class/list", classController.GetClassListHandler)
		r.GET("/class/:id", classController.GetClassByIDHandler)

		// for tags
		r.GET("/course/tag", tagController.GetTagListHandler)
		r.GET("/course/:id/tag", tagController.GetTagByCourseIDHandler)
		r.POST("/course/:id/tag", tagController.CreateTagByCourseIDHandler)
		r.PUT("/course/:id/tag", tagController.VoteTagHandler)

		// for professors
		r.GET("/professors", ProfessorController.GetProfessorListHandler)

		// for testing
		r.GET("/ping", auth.JWTAuth(), func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Content not found",
		})
	})
	return r
}
