package request

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLoggedIn = errors.New("user not logged in")
var ErrorUserUnauthorized = errors.New("user not authorized")

// GetCurrentUser gets the ID of the current user on context
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get("user_ID")
	if !ok {
		err = ErrorUserNotLoggedIn
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLoggedIn
		return
	}
	return
}

// GetCurrentUserRole gets the user role of the current user for authorization
func GetCurrentUserRole(c *gin.Context) (role int32, err error) {
	r, ok := c.Get("role")
	if !ok {
		err = ErrorUserUnauthorized
		return
	}
	role, ok = r.(int32)
	if !ok {
		err = ErrorUserUnauthorized
		return
	}
	return
}
