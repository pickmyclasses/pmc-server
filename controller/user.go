package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// RegisterHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var params model.RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	if err := logic.Register(&params); err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

// LoginHandler User login interface
// @Summary Use this API to login
// @Description After login, a token will be returned to verify user in the future
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.LoginParams true "login parameters"
// @Success 200 {string} jwt token
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var params model.LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	userInfo, err := logic.Login(&params)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: userInfo,
	})
}

func GetUserHistoryHandler(c *gin.Context) {
	userID := c.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	history, err := logic.GetUserHistoryCourseList(int64(userIDInt))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: history,
	})
}

type HistoryParam struct {
	CourseID int64 `json:"courseID"`
	UserID int64 `json:"userID"`
}

func AddUserHistoryHandler(c *gin.Context) {
	var param HistoryParam
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.AddUserCourseHistory(param.UserID, param.CourseID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

func RemoveUserHistoryHandler(c *gin.Context) {
	var param HistoryParam
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.RemoveUserCourseHistory(param.UserID, param.CourseID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
