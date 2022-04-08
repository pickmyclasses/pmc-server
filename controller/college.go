package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/shared"

	"github.com/gin-gonic/gin"
)

// GetCollegeListHandler API for getting the list of the college we have in database
// @Summary Use this API to fetch the entire list of colleges
// @Description This API will only be used for registration
// @Tags Admin
// @Accept application/json
// @Produce application/json
// @Success 200 {string} ok
// @Router /admin/user/list [get]
func GetCollegeListHandler(ctx *gin.Context) {
	collegeList, err := logic.GetCollegeList()
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: collegeList,
	})
}

func GetCollegeByIDHandler(ctx *gin.Context) {

}

func GetCollegeBuildingsHandler(ctx *gin.Context) {

}

func GetCollegeBuildingByIDHandler(ctx *gin.Context) {

}

func GetCollegeSemesterListHandler(ctx *gin.Context) {
	var collegeID int
	var err error
	if collegeID, err = strconv.Atoi(ctx.Param("id")); err != nil || collegeID == 0 {
		_ = ctx.Error(shared.ParamIncompatibleErr{})
		return
	}

	semesterList, err := logic.GetCollegeSemesterList(int32(collegeID))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		shared.DATA: semesterList,
	})
}
