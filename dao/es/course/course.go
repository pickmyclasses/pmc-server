package course

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"pmc_server/init/es"
	esModel "pmc_server/model/es"
	"pmc_server/shared"
	"unicode"
)

// BoolQuery represents the entity of a boolean query in Elasticsearch
type BoolQuery struct {
	index      string
	query      *elastic.BoolQuery
	pageNumber int
	pageSize   int
	context    context.Context
}

// NewBoolQuery creates a new BoolQuery object
func NewBoolQuery(pageNum, pageSize int) *BoolQuery {
	if pageNum < 0 {
		pageNum = 0
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
	// change the weights when there is only string
	// something like cs 5000 should have heavier weights on title than description
	fields := []string{"title^1.5", "description^1.0", "designation_catalog^1.5", "catalog_course_name^1.5"}
	for _, s := range keywords {
		if unicode.IsDigit(s) {
			fields = []string{"title4.0", "description^1.0",
				"designation_catalog^2.0", "catalog_course_name^3.0"}
		}
	}
	c.query = c.query.
		Must(elastic.NewMultiMatchQuery(keywords,
			fields...).
			Fuzziness("AUTO"))
}

func (c *BoolQuery) QueryByMinCredit(minCredit float32) {
	c.query = c.query.Filter(elastic.NewRangeQuery("min_credit").Gte(minCredit))
}

func (c *BoolQuery) QueryByMaxCredit(maxCredit float32) {
	c.query = c.query.Filter(elastic.NewRangeQuery("max_credit").Lte(maxCredit))
}

func (c *BoolQuery) QueryByIsHonor(isHonor bool) {
	c.query = c.query.Filter(elastic.NewTermsQuery("is_honor", isHonor))
}

func (c *BoolQuery) QueryByTypes(types string) {
	c.query = c.query.Filter(elastic.NewMatchQuery("subject", types))
}

func (c *BoolQuery) QueryByOffering() {
	c.query = c.query.Filter(elastic.NewMatchQuery("hasOffering", true))
}

func (c *BoolQuery) QueryByWeekdays(weekdays string) {
	c.query = c.query.Filter(elastic.NewNestedQuery("classes",
		elastic.NewMatchQuery("classes.offer_date", weekdays)))
}

func (c *BoolQuery) QueryByStartTime(startTime float32) {
	c.query = c.query.Filter(elastic.NewNestedQuery("classes",
		elastic.NewRangeQuery("classes.start_time").Gte(startTime)))
}

func (c *BoolQuery) QueryByEndTime(endTime float32) {
	c.query = c.query.Filter(elastic.NewNestedQuery("classes",
		elastic.NewRangeQuery("classes.end_time").Lte(endTime)))
}

func (c *BoolQuery) QueryByProfessors(professors []string) {
	c.query = c.query.Must(elastic.NewNestedQuery("classes",
		elastic.NewTermsQuery("classes.instructors", professors)))
}

func (c *BoolQuery) QueryByTags(tags []string) {
	c.query = c.query.Must(elastic.NewTermQuery("tags", tags))
}

func (c *BoolQuery) QueryByRating(rating float32) {
	c.query = c.query.Filter(elastic.NewRangeQuery("rating").Gte(rating))
}

func (c *BoolQuery) QueryByOnline() {
	c.query = c.query.Filter(elastic.NewNestedQuery("classes",
		elastic.NewTermsQuery("classes.start_time", 0)))
}

func (c *BoolQuery) DoSearch() (*[]int64, int64, error) {
	res, err := es.Client.Search().
		Index(c.index).
		Query(c.query).
		From(c.pageNumber).
		Size(c.pageSize).
		Do(c.context)

	if err != nil {
		return nil, -1, shared.InternalErr{}
	}

	var esCourseIdList []int64
	total := res.Hits.TotalHits.Value

	for _, hit := range res.Hits.Hits {
		var course esModel.Course
		err := json.Unmarshal(*&hit.Source, &course)
		if err != nil {
			return nil, -1, shared.InternalErr{}
		}
		esCourseIdList = append(esCourseIdList, course.ID)
	}

	return &esCourseIdList, total, nil
}
