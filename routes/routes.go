package routes

import (
	"net/http"
	"pmc_server/controller/admin"
	"pmc_server/init/logger"

	classController "pmc_server/controller/class"
	courseController "pmc_server/controller/course"
	reviewController "pmc_server/controller/review"
	scheduleController "pmc_server/controller/schedule"
	userController "pmc_server/controller/user"
	_ "pmc_server/docs"
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

	// for user
	r.POST("/register", userController.RegisterHandler).Use(auth.Cors())
	r.POST("/login", userController.LoginHandler).Use(auth.Cors())

	// for schedule
	r.POST("/:id/schedule", scheduleController.AddUserScheduleHandler).Use(auth.Cors())
	r.GET("/:id/schedule", scheduleController.GetUserScheduleHandler).Use(auth.Cors())
	r.DELETE("/:id/schedule", scheduleController.DeleteUserScheduleHandler).Use(auth.Cors())

	// for course
	r.GET("/course/list", courseController.GetCourseListHandler).Use(auth.Cors())
	r.GET("/course/:id", courseController.GetCourseByIDHandler).Use(auth.Cors())
	r.GET("/course/:id/class", courseController.GetClassesOfCourseHandler).Use(auth.Cors())

	// for course search
	r.POST("/course/search", courseController.GetCoursesBySearchHandler).Use(auth.Cors())

	// for review
	r.GET("/course/:id/review", reviewController.GetCourseReviewListHandler).Use(auth.Cors())
	r.GET("/course/:id/review/:review_id", reviewController.GetCourseReviewByIDHandler).Use(auth.Cors())
	r.POST("/course/review", reviewController.PostCourseReviewHandler).Use(auth.Cors())
	r.PUT("/course/review", reviewController.UpdateCourseReviewHandler).Use(auth.Cors())

	// for class
	r.GET("/class/list", classController.GetClassListHandler).Use(auth.Cors())
	r.GET("/class/:id", classController.GetClassByIDHandler).Use(auth.Cors())

	// for testing
	r.GET("/ping", auth.JWTAuth(), auth.Cors(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// for swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// for admin
	r.GET("/admin/user/list", auth.IsAdminAuth(), admin.GetUserListHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Content not found",
		})
	})
	return r
}
