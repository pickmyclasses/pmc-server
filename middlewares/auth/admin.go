package auth

import (
	"net/http"
	"pmc_server/middlewares/request"

	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// IsAdminAuth is used for checking if the logged-in user is an admin
func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentUser, err := request.GetCurrentUserRole(ctx)
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
