package review

import (
	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"
	"pmc_server/utils"

	"github.com/gin-gonic/gin"
)

func GetCourseReviewListHandler(c *gin.Context) {
	pnInt, pSizeInt, err := utils.HandlePagination(c, "20")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: err,
		})
		return
	}

	courseID := c.Param("id")
	reviewList, total, err := logic.GetCourseReviewList(pnInt, pSizeInt, courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			MESSAGE: NO_INFO_ERR,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    reviewList,
		TOTAL:   total,
	})
}

func GetCourseReviewByIDHandler(c *gin.Context) {
	reviewID := c.Param("review_id")
	if reviewID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: NO_ID_ERR,
		})
	}

	review, err := logic.GetReviewByID(reviewID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			MESSAGE: err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		DATA:    review,
	})
}

func PostCourseReviewHandler(c *gin.Context) {

}
