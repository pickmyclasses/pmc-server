package auth

import (
	"net/http"

	. "pmc_server/consts"
	"pmc_server/controller"

	"github.com/gin-gonic/gin"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentUser, err := controller.GetCurrentUserRole(ctx)
		if err != nil {
			ctx.Abort()
		}

		if currentUser != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				MESSAGE: "Not authorized",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
