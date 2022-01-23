package routes

import (
	"net/http"
	"pmc_server/controller/admin"
	"pmc_server/init/logger"

	classController "pmc_server/controller/class"
	courseController "pmc_server/controller/course"
	reviewController "pmc_server/controller/review"
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
	r.POST("/register", userController.RegisterHandler)
	r.POST("/login", userController.LoginHandler)

	// for course
	r.GET("/course/list", courseController.GetCourseListHandler)
	r.GET("/course/:id", courseController.GetCourseByIDHandler)
	r.GET("/course/:id/class", courseController.GetClassesOfCourseHandler)
	r.GET("/course/:id/review", reviewController.GetCourseReviewListHandler)
	r.GET("/course/:id/review/:review_id", reviewController.GetCourseReviewByIDHandler)
	r.POST("/course/:id/review", reviewController.PostCourseReviewHandler)

	// for class
	r.GET("/class/list", classController.GetClassListHandler)
	r.GET("/class/:id", classController.GetClassByIDHandler)

	// for testing
	r.GET("/ping", auth.JWTAuth(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// for swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// for admin
	r.GET("/admin/user/list", admin.GetUserListHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Content not found",
		})
	})
	return r
}
