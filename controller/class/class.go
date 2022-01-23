package controller

import (
	"net/http"
	"strconv"

	. "pmc_server/consts"
	"pmc_server/logic"
	model "pmc_server/model"

	"github.com/gin-gonic/gin"
)

func GetClassListHandler(c *gin.Context) {
	pn := c.DefaultQuery("pn", "0")
	pSize := c.DefaultQuery("psize", "20")
	pnInt, err := strconv.Atoi(pn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: BAD_PAGE_NUMBER_ERR,
		})
		return
	}
	pSizeInt, err := strconv.Atoi(pSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: BAD_PAGE_SIZE_ERR,
		})
		return
	}

	if pnInt < 0 || pSizeInt < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: BAD_PAGE_ERR,
		})
		return
	}

	classList, total := logic.GetClassList(pnInt, pSizeInt)
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    classList,
		TOTAL:   total,
	})

}

func GetClassByIDHandler(c *gin.Context) {
	var classParam model.ClassParams
	if err := c.ShouldBindUri(&classParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: NO_ID_ERR,
		})
		return
	}

	classInfo, err := logic.GetClassByID(classParam.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			MESSAGE: err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    classInfo,
	})
}
