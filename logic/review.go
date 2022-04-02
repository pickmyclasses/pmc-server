package logic

import (
	"fmt"
	"strconv"

	classDao "pmc_server/dao/postgres/class"
	collegeDao "pmc_server/dao/postgres/college"
	historyDao "pmc_server/dao/postgres/history"
	reviewDao "pmc_server/dao/postgres/review"
	semesterDao "pmc_server/dao/postgres/semester"
	userDao "pmc_server/dao/postgres/user"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
)

func GetCourseReviewList(pn, pSize int, courseID string) (*dto.ReviewList, error) {
	idInt, err := strconv.Atoi(courseID)
	if err != nil {
		return nil, shared.MalformedIDErr{}
	}

	rating, err := reviewDao.GetCourseOverallRating(int64(idInt))
	if err != nil {
		return nil, err
	}

	total, err := reviewDao.GetReviewTotalByCourseID(idInt)
	if err != nil {
		return nil, err
	}

	reviewList, err := reviewDao.GetReviewsByCourseID(idInt, pn, pSize)
	if err != nil {
		return nil, err
	}

	reviewRsp := &dto.ReviewList{
		Total:         total,
		Reviews:       make([]dto.Review, 0),
		OverallRating: rating.OverAllRating,
	}

	for _, review := range reviewList {
		var username string
		userInfo, err := userDao.GetUserByID(review.UserID)
		if err != nil || review.Anonymous {
			username = "Anonymous student"
		} else {
			username = fmt.Sprintf("%s %s", userInfo.FirstName, userInfo.LastName)
		}
		reviewDto := dto.Review{
			Rating:             review.Rating,
			Anonymous:          review.Anonymous,
			Recommended:        review.Recommended,
			Pros:               review.Pros,
			Cons:               review.Cons,
			Comment:            review.Comment,
			CourseID:           int64(idInt),
			UserID:             review.UserID,
			Username:           username,
			CreatedAt:          review.CreatedAt,
			LikedCount:         review.LikeCount,
			DislikedCount:      review.DislikeCount,
			HourSpent:          review.HourSpent,
			GradeReceived:      review.GradeReceived,
			IsExamHeavy:        review.ExamHeavy,
			IsHomeworkHeavy:    review.HomeworkHeavy,
			ExtraCreditOffered: review.ExtraCreditOffered,
		}

		userCourseHistory, err := historyDao.GetUserCourseHistoryByID(review.UserID, review.CourseID)
		if err != nil {
			return nil, err
		}

		semester, err := semesterDao.GetSemesterByID(userCourseHistory.SemesterID)
		if err != nil {
			return nil, err
		}

		college, err := collegeDao.GetCollegeByID(semester.CollegeID)
		if err != nil {
			return nil, err
		}
		reviewDto.ClassSemester = dto.Semester{
			CollegeName: college.Name,
			Year:        semester.Year,
			Season:      semester.Season,
		}

		class, err := classDao.GetClassInfoByID(int(userCourseHistory.ClassID))
		if err != nil {
			return nil, err
		}

		reviewDto.ClassProfessor = class.Instructors
		reviewRsp.Reviews = append(reviewRsp.Reviews, reviewDto)
	}

	return reviewRsp, nil

}

func GetReviewByID(reviewID string) (*model.Review, error) {
	idInt, err := strconv.Atoi(reviewID)
	if err != nil {
		return nil, shared.MalformedIDErr{}
	}
	return reviewDao.GetReviewByID(idInt)
}

func PostCourseReview(review dto.Review, courseID int64) error {
	userCourseHistory, err := historyDao.CheckIfCourseInUserCourseHistory(review.UserID, courseID)
	if err != nil {
		return err
	}
	if !userCourseHistory {
		return shared.NoPreviousRecordErr{}
	}
	// check old overall rating value
	rating, err := reviewDao.GetCourseOverallRating(courseID)
	if err != nil {
		return err
	}

	// if there is no existing rating, create a new one first
	if rating == nil || rating.ID == 0 {
		rating, err = reviewDao.CreateCourseRating(courseID)
		if err != nil {
			return err
		}
	}

	// recalculate the rating for the record
	rating.OverAllRating =
		((rating.OverAllRating * float32(rating.TotalRatingCount)) + review.Rating) /
			(float32(rating.TotalRatingCount) + 1)

	reviewRec := &model.Review{
		Rating:             review.Rating,
		Anonymous:          review.Anonymous,
		Recommended:        review.Recommended,
		Pros:               review.Pros,
		Cons:               review.Cons,
		Comment:            review.Comment,
		CourseID:           courseID,
		UserID:             review.UserID,
		LikeCount:          0,
		DislikeCount:       0,
		HourSpent:          review.HourSpent,
		GradeReceived:      review.GradeReceived,
		ExamHeavy:          review.IsExamHeavy,
		HomeworkHeavy:      review.IsHomeworkHeavy,
		ExtraCreditOffered: review.ExtraCreditOffered,
	}

	err = reviewDao.CreateCourseReview(*reviewRec)
	if err != nil {
		return err
	}

	rating.TotalRatingCount += 1

	// save the new rating
	err = reviewDao.UpdateCourseRating(*rating)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCourseReview(review model.ReviewParams) error {
	return reviewDao.UpdateCourseReview(review)
}
