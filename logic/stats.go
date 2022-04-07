package logic

import (
	courseDao "pmc_server/dao/postgres/course"
	historyDao "pmc_server/dao/postgres/history"
	dao "pmc_server/dao/postgres/review"
	"pmc_server/shared"
	"sort"
)

type ProfessorRanking struct {
	Name   string  `json:"name"`
	Rating float32 `json:"rating"`
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
				mapping[history.ProfessorName] = struct {
					Rating float32
					Count  int32
				}{Rating: mapping[history.ProfessorName].Rating, Count: count}
			} else {
				mapping[history.ProfessorName] = struct {
					Rating float32
					Count  int32
				}{Rating: review.Rating, Count: 1}
			}
		}
	}

	profRankingList := make([]ProfessorRanking, 0)
	for k, v := range mapping {
		profRank := ProfessorRanking{
			Name:   k,
			Rating: v.Rating / float32(v.Count),
		}
		profRankingList = append(profRankingList, profRank)
	}

	// sort the list from high to low rating
	sort.Slice(profRankingList, func(i, j int) bool {
		return profRankingList[i].Rating > profRankingList[j].Rating
	})

	return profRankingList, nil
}

type CourseLoad struct {
	CourseAverageGrade float32 `json:"courseAverageGrade"`
	MajorAverageGrade  float32 `json:"majorAverageGrade"`
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

func getCourseNumberGrade(grade string) (float32, error) {
	mapping := map[string]float32{
		"A":  4.0,
		"A-": 3.7,
		"B+": 3.3,
		"B":  3.0,
		"B-": 2.7,
		"C+": 2.3,
		"C":  2.0,
		"C-": 1.7,
		"D+": 1.3,
		"D":  1.0,
		"E":  0.0,
		"F":  0.0,
	}
	if _, ok := mapping[grade]; ok {
		return mapping[grade], nil
	}
	return 0, shared.InfoUnmatchedErr{}
}
