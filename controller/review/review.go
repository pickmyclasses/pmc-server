package review

import (
	"net/http"
	"pmc_server/model/dto"

	. "pmc_server/consts"
	"pmc_server/logic"
	"pmc_server/utils"

	"github.com/gin-gonic/gin"
)

// GetCourseReviewListHandler Get review list for the given course
// @Summary Use this API to get the list of reviews of the given course
// @Description This API is used to get the review list of the given course, you should do pagination
// @Tags Review
// @Accept application/json
// @Produce application/json
// @Param id query int true "course ID "
// @Param pn query int false "Page number of the paginated review list, default is 1" mininum(1)
// @Param psize query int false "Page size(number of reviews to fetch) of the paginated review list, default is 10" mininum(1) maxinum(30)
// @Success 200 {int} total number of the reviews
// @Success 200 {array} dto.Review
// @Router /course/:id/review [get]
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

// GetCourseReviewByIDHandler Get a single review under a course by giving its ID
// @Summary Use this API to get a single review under a class
// @Description This API is used when user click on a single review and gets its detail, I will say it's hardly useful
// @Tags Review
// @Accept application/json
// @Produce application/json
// @Param id query int true "The course id which the review is posted to"
// @Param review_id query int true "The actual review ID"
// @Success 200 {object} dto.Review
// @Router /course/:id/review/:review_id [get]
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

// PostCourseReviewHandler Post a single review under a course
// @Summary Use this API to post a single review under a course
// @Description This API is for posting a review under a course
// @Tags Review
// @Accept application/json
// @Produce application/json
// @Param object body model.ReviewParams true "Post review parameters"
// @Success 200 {string} OK
// @Router /course/review [post]
func PostCourseReviewHandler(c *gin.Context) {
	var review dto.Review
	if err := c.ShouldBind(&review); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			MESSAGE: err.Error(),
		})
		return
	}

	err := logic.PostCourseReview(review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
