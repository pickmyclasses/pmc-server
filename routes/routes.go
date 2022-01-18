package routes

import (
	"net/http"

	classController "pmc_server/controller/class"
	courseController "pmc_server/controller/course"
	reviewController "pmc_server/controller/review"
	userController "pmc_server/controller/user"
	"pmc_server/logger"
	"pmc_server/middlewares/auth"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/register", userController.RegisterHandler)
	r.POST("/login", userController.LoginHandler)

	r.GET("/course/list", courseController.GetCourseListHandler)
	r.GET("/course/:id", courseController.GetCourseHandler)
	r.GET("/course/class/:id", courseController.GetClassesOfCourseHandler)
	r.GET("/course/:id/review", reviewController.GetCourseReviewListHandler)
	r.GET("/course/:id/review/:review_id", reviewController.GetCourseReviewByIDHandler)
	r.POST("/course/:id/review", reviewController.PostCourseReviewHandler)

	r.GET("/class/list", classController.GetClassListHandler)
	r.GET("/class/:id", classController.GetClassByIDHandler)

	// for testing
	r.GET("/ping", auth.JWTAuth(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/admin/user/list", userController.GetUserListHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Content not found",
		})
	})
	return r
}
