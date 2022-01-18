package controller

import (
	"errors"

	"github.com/gin-gonic/gin"

	"pmc_server/middlewares/auth"
)

var ErrorUserNotLoggedIn = errors.New("user not logged in")

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(auth.CtxUserIDKey)
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
