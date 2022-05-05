// Package controller - controller for the class entities
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

// GetClassListHandler Get the entire class list
// @Summary Use this API to get the list of the classes
// @Description This API is used to get the class list of the given course, you should do pagination
// @Tags Class
// @Accept application/json
// @Produce application/json
// @Param pn query int false "Page number of the paginated class list, default is 1" mininum(1)
// @Param psize query int false "Page size(number of class to fetch) of the paginated class list, default is 10" mininum(1) maxinum(30)
// @Success 200 {int} total number of the classes
// @Success 200 {array} dto.Class
// @Router /class/list [get]
func GetClassListHandler(c *gin.Context) {
	pn := c.DefaultQuery("pn", "0")
	pSize := c.DefaultQuery("psize", "20")
	pnInt, err := strconv.Atoi(pn)

	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	pSizeInt, err := strconv.Atoi(pSize)
	if err != nil {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	if pnInt < 0 || pSizeInt < 0 {
		_ = c.Error(shared.ParamIncompatibleErr{})
		return
	}

	classList, total := logic.GetClassList(pnInt, pSizeInt)
	c.JSON(http.StatusOK, gin.H{
		DATA:  classList,
		TOTAL: total,
	})

}

// GetClassByIDHandler Get class by the given ID
// @Summary Use this API to get the class by the given ID
// @Description This API is used to get the class by the given ID
// @Tags Review
// @Accept application/json
// @Produce application/json
// @Param id query int true "class id"
// @Success 200 {object} dto.Class
// @Router /class/:id [get]
func GetClassByIDHandler(c *gin.Context) {
	var classParam model.ClassParams
	if err := c.ShouldBindUri(&classParam); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	classInfo, err := logic.GetClassByID(classParam.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		DATA: classInfo,
	})
}
