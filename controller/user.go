// Package controller - controller for user entities
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
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

// PostUserMajorParams defines the parameters for users to post their major
type PostUserMajorParams struct {
	UserID     int64  `json:"userID"`       // the ID of the user
	MajorName  string `json:"majorName"`    // the name of the major of the user
	Emphasis   string `json:"emphasisName"` // the name of the emphasis of the user, this can be empty
	SchoolYear string `json:"schoolYear"`   // the school year of the user, i.e. senior, junior and such
}

// HistoryParam defines the parameters for users to fetch/post their course history
type HistoryParam struct {
	CourseID int64 `json:"courseID"` // the ID of the course user has taken
	UserID   int64 `json:"userID"`   // the ID of the user
}

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

// GetUserHistoryHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
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

// AddUserHistoryHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
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

// RemoveUserHistoryHandler User registration interface
// @Summary Use this API to register a user
// @Description You should only use this interface to register for student, professor/admin should be assigned directly
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body model.RegisterParams true "registration parameters"
// @Success 200 {string} success
// @Router /register [post]
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

// PostUserMajorHandler posts the major of the user to the database
// @Summary Use this API to post user's major info including emphasis and year info
// @Description this API is used to post major/emphasis info of the user, emphasis could be empty
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param object body PostUserMajorParams true "Post major/emphasis/year parameters"
// @Success 200 {string} success
// @Router /user/major [post]
func PostUserMajorHandler(c *gin.Context) {
	var param PostUserMajorParams
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	user, err := logic.PostUserMajor(param.UserID, param.MajorName, param.Emphasis, param.SchoolYear)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    user,
	})
}

// GetUserRecommendCourseHandler fetches the recommended course list of the given user
// @Summary Use this API to get the recommended course list of the given user along with the degree catalog
// @Description This API will give the degree catalogs of the user and major, 8 top courses for each degree catalog
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param int parameter int true "user ID"
// @Success 200 {string} success
// @Router /register [post]
func GetUserRecommendCourseHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{
			Msg: "Unable to process the given user ID",
		})
	}

	recommendCourses, err := logic.RecommendCourses(int64(uid))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: recommendCourses,
	})
}

func GetUserBookmark(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)

	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{
			Msg: "Unable to process the given user ID",
		})
	}

	bookmarkList, err := logic.GetUserBookmarks(int64(uid))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		DATA: bookmarkList,
	})
}

type BookmarkParam struct {
	CourseID int64 `json:"courseID"`
}

func AddUserBookmark(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)

	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{
			Msg: "Unable to process the given user ID",
		})
	}

	var param BookmarkParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		_ = ctx.Error(shared.ParamInsufficientErr{})
		return
	}

	err = logic.PostUserBookmark(int64(uid), param.CourseID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

func DeleteUserBookmark(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)

	if err != nil {
		_ = ctx.Error(shared.ParamIncompatibleErr{
			Msg: "Unable to process the given user ID",
		})
	}

	var param BookmarkParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		_ = ctx.Error(shared.ParamInsufficientErr{})
		return
	}

	err = logic.DeleteUserBookmark(int64(uid), param.CourseID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
