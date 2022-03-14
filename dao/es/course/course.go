package course

import (
	"context"
	"encoding/json"
	"pmc_server/init/es"
	esModel "pmc_server/model/es"
	"pmc_server/shared"
	. "pmc_server/shared"

	"github.com/olivere/elastic/v7"
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
	c.query = c.query.
		Must(elastic.NewMultiMatchQuery(keywords,
			"title", "description", "designation_catalog", "catalog_course_name").Fuzziness("AUTO"))
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
