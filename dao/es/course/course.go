package course

import (
	"context"
	. "pmc_server/consts"
	"pmc_server/init/es"

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

func (c *BoolQuery) DoSearch() (*elastic.SearchResult, error) {
	return es.Client.Search().Index(c.index).Query(c.query).From(c.pageNumber).Size(c.pageSize).Do(c.context)
}
