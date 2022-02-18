package course

import (
	"errors"

	"pmc_server/init/postgres"
	"pmc_server/model"
	"pmc_server/utils"

	"github.com/olivere/elastic/v7"
)

func GetCourses(pn, pSize int) ([]model.Course, int64) {
	var courses []model.Course
	total := postgres.DB.Find(&courses).RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Find(&courses)

	return courses, total
}

func GetCourseByID(id int) (*model.Course, error) {
	var course model.Course
	result := postgres.DB.Where("id = ?", id).First(&course)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no course info found")
	}
	return &course, nil
}

func GetClassListByCourseID(id int) (*[]model.Class, int64) {
	var classes []model.Class
	res := postgres.DB.Where("course_id = ?", id).Find(&classes)
	return &classes, res.RowsAffected
}

func GetCoursesBySearch(param model.CourseFilterParams) (*[]model.Course, int32, error) {
	query := elastic.NewBoolQuery()
	if param.MinCredit != 0 {
		query = query.Filter(elastic.NewRangeQuery("min_credit").Gte(param.MinCredit))
	}

	if param.MaxCredit != 0 {
		query = query.Filter(elastic.NewRangeQuery("max_credit").Lte(param.MaxCredit))
	}

	if param.StartTime != 0 {
		query = query.Filter(elastic.NewRangeQuery("start_time").Gte(param.StartTime))
	}

	if param.EndTime != 0 {
		query = query.Filter(elastic.NewRangeQuery("end_time").Lte(param.EndTime))
	}

	if param.IsHonor {
		query = query.Must(elastic.NewTermsQuery("is_honor", true))
	}

	if param.OfferedOnline {
	}

	if param.OfferedOffline {

	}

	if param.RankByRatingHighToLow {

	}

	if param.RankByRatingLowToHigh {

	}

	if param.MinRating != 0 {

	}

	if len(param.Weekday) != 0 {

	}

	if param.Keyword != "" {
		query = query.Must(elastic.NewMultiMatchQuery(param.Keyword, "title", "description", "designation_catalog", "catalog_course_name"))
	}

	return nil, -1, nil
}
