package class

import (
	"context"
	"encoding/json"
	"pmc_server/init/es"
	esModel "pmc_server/model/es"
	"pmc_server/shared"

	"github.com/olivere/elastic/v7"
)

type BoolQuery struct {
	index      string
	query      *elastic.BoolQuery
	context    context.Context
	pageSize   int
	pageNumber int
}

func NewBoolQuery(pageSize int) *BoolQuery {
	return &BoolQuery{
		query:      elastic.NewBoolQuery(),
		index:      "class",
		context:    context.Background(),
		pageSize:   pageSize,
		pageNumber: 0,
	}
}

func (c *BoolQuery) QueryByIsOnline(isOnline bool) {
	c.query = c.query.Must(elastic.NewTermsQuery("is_online", isOnline))
}

func (c *BoolQuery) QueryByIsInPerson(isInPerson bool) {
	c.query = c.query.Must(elastic.NewTermsQuery("is_in_person", isInPerson))
}

func (c *BoolQuery) QueryByIsIVC(isIVC bool) {
	c.query = c.query.Must(elastic.NewTermsQuery("is_ivc", isIVC))
}

func (c *BoolQuery) QueryByIsHybrid(isHybrid bool) {
	c.query = c.query.Must(elastic.NewTermsQuery("is_hybrid", isHybrid))
}

func (c *BoolQuery) QueryByOfferDates(offerDates []int) {
	status := make([]interface{}, len(offerDates))
	for i, v := range offerDates {
		status[i] = v
	}
	c.query = c.query.Filter(elastic.NewTermsQuery("offer_dates", status...))
}

func (c *BoolQuery) QueryByOfferTime(startTime, endTime float32) {
	c.query = c.query.
		Must(elastic.NewRangeQuery("start_time").Gte(startTime)).
		Must(elastic.NewRangeQuery("end_time").Lte(endTime))
}

func (c *BoolQuery) QueryByProfessor(professorName string) {
	c.query = c.query.Must(elastic.NewFuzzyQuery(professorName, "professors").Fuzziness("AUTO"))
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

	var esCourseIDList []int64
	total := res.Hits.TotalHits.Value

	for _, hit := range res.Hits.Hits {
		var course esModel.Class
		err := json.Unmarshal(*&hit.Source, &course)
		if err != nil {
			return nil, -1, shared.InternalErr{}
		}
		esCourseIDList = append(esCourseIDList, course.ID)
	}

	return &esCourseIDList, total, nil
}
