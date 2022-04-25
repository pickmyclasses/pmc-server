package logic

import (
	"encoding/json"
	"fmt"
	"net/http"
	classDao "pmc_server/dao/postgres/class"
	collegeDao "pmc_server/dao/postgres/college"
	courseDao "pmc_server/dao/postgres/course"
	historyDao "pmc_server/dao/postgres/history"
	reviewDao "pmc_server/dao/postgres/review"
	dao "pmc_server/dao/postgres/schedule"
	semesterDao "pmc_server/dao/postgres/semester"
	tagDao "pmc_server/dao/postgres/tag"
	userDao "pmc_server/dao/postgres/user"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
	"strconv"
	"strings"
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

	reviewList, err := reviewDao.GetPaginatedReviewsByCourseID(idInt, pn, pSize)
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

		if err != nil {
			return nil, err
		}

		reviewDto.ClassProfessor = userCourseHistory.ProfessorName
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

func PostCourseReview(review dto.Review, courseID int64, extraInfoNeeded bool) error { // analyze keywords
	if extraInfoNeeded {
		semester, err := semesterDao.GetSemesterByName(review.ClassSemester.Season, review.ClassSemester.Year)
		if err != nil {
			return err
		}

		exist, err := historyDao.CheckIfCourseInUserCourseHistory(review.UserID, courseID)
		if err != nil {
			return err
		}
		if exist {
			err = historyDao.UpdateHistoryByCourseID(review.UserID, review.CourseID, int32(semester.ID), review.ClassProfessor)
			if err != nil {
				return err
			}
		} else {
			err = historyDao.CreateSingleUserCourseHistory(review.UserID, courseID, int32(semester.ID),
				review.ClassProfessor)
			if err != nil {
				return err
			}
		}
	}

	userCourseHistoryExist, err := historyDao.CheckIfCourseInUserCourseHistory(review.UserID, courseID)
	if err != nil {
		return err
	}
	if !userCourseHistoryExist {
		return shared.NoPreviousRecordErr{}
	}
	history, err := historyDao.GetUserCourseHistoryByID(review.UserID, courseID)
	if err != nil {
		return err
	}
	if history.ProfessorName == "" || history.SemesterID == 0 {
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

	tagList, err := analyzeTags(review.Tags)
	if err != nil {
		return err
	}

	// insert tags
	for _, tag := range tagList {
		err = tagDao.CreateTagByCourseID(courseID, strings.TrimSpace(tag.Content), int32(tag.Type))
		if err != nil {
			return err
		}
	}

	reviewRec := &model.Review{
		Rating:             review.Rating,
		Anonymous:          review.Anonymous,
		Recommended:        review.Recommended,
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
		Tags:               review.Tags,
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

func VoteCourseReview(userID, courseID, voterID int64, isUpvote bool) error {
	return nil
}

type UserReviewInfo struct {
	HasTaken        bool       `json:"hasTaken"`
	HasReviewed     bool       `json:"hasReviewed"`
	CurrentlyTaking bool       `json:"currentlyTaking"`
	ReviewContent   dto.Review `json:"reviewContent"`
}

func GetUserReviewInfo(userID, courseID int64) (*UserReviewInfo, error) {
	info := &UserReviewInfo{}
	// this logic is not perfect, fix it later by removing all the scheduled items after the semester ends
	scheduleList, err := dao.GetScheduleByUserID(userID)
	if err != nil {
		return nil, err
	}

	currentlyTaking := false
	for _, s := range scheduleList {
		class, err := classDao.GetClassByID(int(s.ClassID))
		if err != nil {
			return nil, err
		}
		course, err := courseDao.GetCourseByID(int(class.CourseID))
		if err != nil {
			return nil, err
		}
		if courseID == course.ID {
			currentlyTaking = true
		}
	}

	historicalTaken := false
	history, err := historyDao.GetUserCourseHistoryByID(userID, courseID)
	if err != nil {
		return nil, err
	}
	if history != nil {
		historicalTaken = true
	}

	hasReviewed := false
	review, err := reviewDao.GetReviewOfUserForACourse(userID, courseID)
	if err != nil {
		return nil, err
	}
	if review != nil {
		hasReviewed = true
	}

	var semester *model.Semester
	if historicalTaken && hasReviewed {
		semester, err = semesterDao.GetSemesterByID(history.SemesterID)
		if err != nil {
			return nil, err
		}
	}
	info.HasTaken = historicalTaken
	info.HasReviewed = hasReviewed
	info.CurrentlyTaking = currentlyTaking
	if info.HasReviewed {
		info.ReviewContent = dto.Review{
			Rating:             review.Rating,
			Anonymous:          review.Anonymous,
			Recommended:        review.Recommended,
			Comment:            review.Comment,
			CourseID:           review.CourseID,
			UserID:             review.UserID,
			CreatedAt:          review.CreatedAt,
			LikedCount:         review.LikeCount,
			DislikedCount:      review.DislikeCount,
			HourSpent:          review.HourSpent,
			GradeReceived:      review.GradeReceived,
			IsExamHeavy:        review.ExamHeavy,
			IsHomeworkHeavy:    review.HomeworkHeavy,
			ExtraCreditOffered: review.ExtraCreditOffered,
			ClassSemester: dto.Semester{
				Year:   semester.Year,
				Season: semester.Season,
			},
			ClassProfessor: history.ProfessorName,
			Tags:           review.Tags,
		}
	}
	return info, nil
}

type Tag struct {
	Content string
	Type    int
}

type TwinwordResp struct {
	ResultMessage string `json:"result_msg"`
	Type          string `json:"type"`
}

func analyzeTags(contents []string) ([]Tag, error) {
	keywordList := make([]Tag, 0)
	for _, content := range contents {
		text := content
		text = strings.ReplaceAll(strings.TrimSpace(text), " ", "%20")
		if text == "" {
			continue
		}

		uri := fmt.Sprintf("https://twinword-sentiment-analysis.p.rapidapi.com/analyze/?text=%s", text)
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Add("X-RapidAPI-Host", "twinword-sentiment-analysis.p.rapidapi.com")
		req.Header.Add("X-RapidAPI-Key", "40604cbd89msh0c3990b01aaabbap141213jsn02ce51189bf0")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()

		var result TwinwordResp
		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			continue
		}
		if result.ResultMessage == "Success" {
			keywordType := -1
			if result.Type == "neutral" || result.Type == "positive" {
				keywordType = 1
			} else {
				keywordType = 0
			}
			keywordList = append(keywordList, Tag{
				Content: strings.TrimSpace(content),
				Type:    keywordType,
			})
		}
	}

	return keywordList, nil
}
