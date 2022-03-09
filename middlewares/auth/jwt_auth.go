package auth

import (
	"net/http"
	"strings"

	"pmc_server/libs/jwt"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "no token provided",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "token format error",
			})
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "token invalid",
			})
			c.Abort()
			return
		}

		c.Set(CtxUserIDKey, claims.UserID)
		c.Set(CtxUserFirstNameKey, claims.FirstName)
		c.Set(CtxUserLastNameKey, claims.LastName)
		c.Set(CtxUserRole, claims.Role)

		c.Next()
	}
}
