package course

import (
	"context"
	"encoding/json"
	"fmt"
	. "pmc_server/consts"
	"pmc_server/init/es"
	esModel "pmc_server/model/es"

	"github.com/olivere/elastic/v7"
)

type BoolQuery struct {
	index      string
	query      *elastic.BoolQuery
	pageNumber int
	pageSize   int
	context    context.Context
}

func NewBoolQuery(pageNum, pageSize int) *BoolQuery {
	if pageNum < 0 {
		pageNum = 0
	}

	if pageSize < 10 {
		pageSize = 10
	}

	return &BoolQuery{
		query:      elastic.NewBoolQuery(),
		index:      "course",
		pageNumber: pageNum,
		pageSize:   pageSize,
		context:    context.Background(),
	}
}

func (c *BoolQuery) QueryByKeywords(keywords string) {
	c.query = c.query.
		Must(elastic.NewMultiMatchQuery(keywords,
			"title", "description", "designation_catalog", "catalog_course_name"))
}

func (c *BoolQuery) QueryByMinCredit(minCredit float32) {
	c.query = c.query.Filter(elastic.NewRangeQuery("min_credit").Gte(minCredit).Lte(MAX))
}

func (c *BoolQuery) QueryByMaxCredit(maxCredit float32) {
	c.query = c.query.Filter(elastic.NewRangeQuery("max_credit").Gte(0).Lte(maxCredit))
}

func (c *BoolQuery) QueryByIsHonor(isHonor bool) {
	c.query = c.query.Must(elastic.NewTermsQuery("is_honor", isHonor))
}

func (c *BoolQuery) QueryByTypes(types string) {
	c.query = c.query.Must(elastic.NewMatchQuery("subject", types))
}

func (c *BoolQuery) DoSearch() ([]int64, int64, error) {
	res, err := es.Client.Search().Index(c.index).Query(c.query).From(c.pageNumber).Size(c.pageSize).Do(c.context)
	if err != nil {
		return nil, -1, fmt.Errorf("error when searching courses %+v", err)
	}

	var esCourseIdList []int64
	total := res.Hits.TotalHits.Value

	for _, hit := range res.Hits.Hits {
		var course esModel.Course
		err := json.Unmarshal(*&hit.Source, &course)
		if err != nil {
			return nil, -1, fmt.Errorf("error when unmarshalling elastic search entity %+v", err)
		}
		esCourseIdList = append(esCourseIdList, course.ID)
	}

	return esCourseIdList, total, nil

	//var courseDtoList []dto.Course
	//for _, id := range esCourseIdList {
	//	course, err := courseDao.GetCourseByID(int(id))
	//	if err != nil {
	//		return nil, -1, fmt.Errorf("error when fetching courses %+v", err)
	//	}
	//
	//	// fetch classes of the course
	//	classList, _ := courseDao.GetClassListByCourseID(int(id))
	//	rating, err := reviewDao.GetCourseOverallRating(id)
	//	if err != nil {
	//		return nil, -1, fmt.Errorf("error when fetching overall rating of the course %+v", err)
	//	}
	//
	//	maxCredit := 0.0
	//	minCredit := 0.0
	//	max, err := strconv.ParseFloat(course.MaxCredit, 32)
	//	if err != nil {
	//		maxCredit = 0.0
	//	}
	//	maxCredit = max
	//
	//	min, err := strconv.ParseFloat(course.MinCredit, 32)
	//	if err != nil {
	//		minCredit = 0.0
	//	}
	//	minCredit = min
	//
	//	courseDto := dto.Course{
	//		CourseID:           id,
	//		IsHonor:            course.IsHonor,
	//		FixedCredit:        course.FixedCredit,
	//		DesignationCatalog: course.DesignationCatalog,
	//		Description:        course.Description,
	//		Prerequisites:      course.Prerequisites,
	//		Title:              course.Title,
	//		CatalogCourseName:  course.CatalogCourseName,
	//		Component:          course.Component,
	//		MaxCredit:          maxCredit,
	//		MinCredit:          minCredit,
	//		Classes:            *classList,
	//		OverallRating:      rating.OverAllRating,
	//	}
	//
	//	courseDtoList = append(courseDtoList, courseDto)
	//}
	//
	//return &courseDtoList, total, nil
}
