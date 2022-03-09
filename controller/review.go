package controller

import (
	"net/http"
	"strconv"

	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
	. "pmc_server/shared"

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
	courseID := c.Param("id")

	pnInt, pSizeInt, err := HandlePagination(c, "20")
	if err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	reviewList, err := logic.GetCourseReviewList(pnInt, pSizeInt, courseID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: reviewList,
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
		_ = c.Error(shared.MalformedIDErr{})
		return
	}

	review, err := logic.GetReviewByID(reviewID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: review,
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
	courseID := c.Param("id")
	courseIDInt, err := strconv.Atoi(courseID)
	if err != nil {
		_ = c.Error(shared.MalformedIDErr{})
		return
	}

	var param dto.Review
	if err := c.ShouldBindJSON(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err = logic.PostCourseReview(param, int64(courseIDInt))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}

// UpdateCourseReviewHandler Updates a single review under a course
// @Summary Use this API to update a single review under a course
// @Description This API is for updating a review under a course
// @Tags Review
// @Accept application/json
// @Produce application/json
// @Param object body model.ReviewParams true "Update review parameters"
// @Success 200 {string} OK
// @Router /course/review [put]
func UpdateCourseReviewHandler(c *gin.Context) {
	var param model.ReviewParams
	if err := c.ShouldBind(&param); err != nil {
		_ = c.Error(shared.ParamInsufficientErr{})
		return
	}

	err := logic.UpdateCourseReview(param)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
	})
}
