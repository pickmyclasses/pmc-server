package logic

import (
	"fmt"
	"sort"

	courseDao "pmc_server/dao/postgres/course"
	historyDao "pmc_server/dao/postgres/history"
	dao "pmc_server/dao/postgres/review"
	semesterDao "pmc_server/dao/postgres/semester"
	"pmc_server/shared"
)

type ProfessorRanking struct {
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
	Grade  float32 `json:"grade"`
}

type SemesterRating struct {
	SemesterName string  `json:"semesterName"`
	Rating       float32 `json:"rating"`
}

type HighestRatedEntity struct {
	TopCourses []struct {
		CourseName string  `json:"courseName"`
		AvgGrade   float32 `json:"avgGrade"`
	} `json:"topCourses"`
	MajorAverage float32 `json:"majorAverage"`
}

type CourseLoad struct {
	CourseAverageGrade float32 `json:"courseAverageGrade"`
	MajorAverageGrade  float32 `json:"majorAverageGrade"`
}

// GetProfessorRankingByCourseID fetches the stats of the average rating each professor received for the course
// since we don't really have individual rating, average score is fetched by different class sessions
func GetProfessorRankingByCourseID(courseID int64) ([]ProfessorRanking, error) {
	reviewList, err := dao.GetReviewListByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	mapping := make(map[string]struct {
		Rating float32
		Count  int32
		Grade  float32
	})
	for _, review := range reviewList {
		history, err := historyDao.GetUserCourseHistoryByID(review.UserID, courseID)
		if err != nil {
			return nil, err
		}
		// this should not happen, but just to be safe (legacy data may cause this to crash)
		if history == nil {
			continue
		}
		// collect record only when it actually has professor provided
		if history.ProfessorName != "" {
			if _, exist := mapping[history.ProfessorName]; exist {
				count := mapping[history.ProfessorName].Count + 1
				rating := mapping[history.ProfessorName].Rating + review.Rating
				newReceived, err := getCourseNumberGrade(review.GradeReceived)
				if err != nil {
					return nil, err
				}
				newGrade := mapping[history.ProfessorName].Grade + newReceived
				mapping[history.ProfessorName] = struct {
					Rating float32
					Count  int32
					Grade  float32
				}{Rating: rating, Count: count, Grade: newGrade}
			} else {
				numberGrade, err := getCourseNumberGrade(review.GradeReceived)
				if err != nil {
					return nil, err
				}
				mapping[history.ProfessorName] = struct {
					Rating float32
					Count  int32
					Grade  float32
				}{Rating: review.Rating, Count: 1, Grade: numberGrade}
			}
		}
	}

	profRankingList := make([]ProfessorRanking, 0)
	for k, v := range mapping {

		profRank := ProfessorRanking{
			Name:   k,
			Rating: v.Rating / float32(v.Count),
			Grade:  v.Grade / float32(v.Count),
		}
		profRankingList = append(profRankingList, profRank)
	}

	// sort the list from high to low rating
	sort.Slice(profRankingList, func(i, j int) bool {
		return profRankingList[i].Rating > profRankingList[j].Rating
	})

	return profRankingList, nil
}

func GetCourseAverageLoad(courseID int64) (*CourseLoad, error) {
	reviewList, err := dao.GetReviewListByCourseID(courseID)
	if err != nil {
		return nil, err
	}
	count := 0
	var gradeTotal float32
	// average score for the course
	if len(reviewList) == 0 {
		return &CourseLoad{
			CourseAverageGrade: 0,
			MajorAverageGrade:  0,
		}, nil
	}
	for _, review := range reviewList {
		count++
		numberGrade, err := getCourseNumberGrade(review.GradeReceived)
		if err != nil {
			return nil, err
		}
		gradeTotal += numberGrade
	}
	courseLoad := gradeTotal / float32(count)

	// average score for the major course (same major as the given course)
	course, err := courseDao.GetCourseByID(int(courseID))
	if err != nil {
		return nil, err
	}

	catalog, _ := shared.ParseString(course.CatalogCourseName, true)
	majorCourseList, err := courseDao.GetCourseListByMajorName(catalog)
	if err != nil {
		return nil, err
	}

	var total float32
	count = 0
	for _, course := range majorCourseList {
		reviewList, err := dao.GetReviewListByCourseID(course.ID)
		if err != nil {
			return nil, err
		}
		if len(reviewList) == 0 {
			continue
		}
		for _, review := range reviewList {
			count++
			numberGrade, err := getCourseNumberGrade(review.GradeReceived)
			if err != nil {
				return nil, err
			}
			total += numberGrade
		}
	}

	majorLoad := total / float32(count)
	return &CourseLoad{
		CourseAverageGrade: courseLoad,
		MajorAverageGrade:  majorLoad,
	}, nil
}

// GetCourseRatingTrendBySemester fetches the overall rating for the course for each semester
// The returned result should be sorted
func GetCourseRatingTrendBySemester(courseID int64) ([]SemesterRating, error) {
	reviewList, err := dao.GetReviewListByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	mapping := make(map[string]struct {
		rating float32
		count  int32
	}, 0)

	for _, review := range reviewList {
		history, err := historyDao.GetUserCourseHistoryByID(review.UserID, courseID)
		if err != nil {
			return nil, err
		}

		semester, err := semesterDao.GetSemesterByID(history.SemesterID)
		if err != nil {
			return nil, err
		}

		semesterName := fmt.Sprintf("%s %d", semester.Season, semester.Year)
		if v, exist := mapping[semesterName]; exist {
			mapping[semesterName] = struct {
				rating float32
				count  int32
			}{rating: v.rating + review.Rating, count: v.count + 1}
		} else {
			mapping[semesterName] = struct {
				rating float32
				count  int32
			}{rating: review.Rating, count: 1}
		}
	}

	semesterRatingList := make([]SemesterRating, 0)
	for k, v := range mapping {
		semesterRatingList = append(semesterRatingList, SemesterRating{
			SemesterName: k,
			Rating:       v.rating / float32(v.count),
		})
	}

	return semesterRatingList, nil
}

type Popularity struct {
	Trends []Trend `json:"trends"`
}

type Trend struct {
	Number   int32  `json:"popularity"`
	Semester string `json:"semester"`
}

func GetCoursePopularity(courseID int64) (*Popularity, error) {
	course, err := courseDao.GetCourseByID(int(courseID))
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, shared.ContentNotFoundErr{
			Msg: fmt.Sprintf("unable to find course %d\n", courseID),
		}
	}

	//TODO: change this college id
	semesterList, err := semesterDao.GetSemesterListByCollegeID(1)
	if err != nil {
		return nil, err
	}

	trendList := make([]Trend, 0)
	for _, s := range semesterList {
		popularity, err := courseDao.GetCoursePopularityBySemesterID(courseID, int32(s.ID))
		if err != nil {
			return nil, err
		}
		if popularity == nil {
			continue
		}

		trendList = append(trendList, Trend{
			Number:   popularity.Popularity,
			Semester: fmt.Sprintf("%s %d", s.Season, s.Year),
		})
	}

	return &Popularity{Trends: trendList}, nil
}
