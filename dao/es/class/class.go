package class

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type BoolQuery struct {
	index   string
	query   *elastic.BoolQuery
	context context.Context
}

func NewBoolQuery() *BoolQuery {
	return &BoolQuery{
		query:   elastic.NewBoolQuery(),
		index:   "course",
		context: context.Background(),
	}
}

func QueryByIsOnline(isOnline bool) {

}

func QueryByIsInPerson(isInPerson bool) {

}

func QueryByIsIVC(isIVC bool) {

}

func QueryByIsHybrid(isHybrid bool) {

}

func QueryByOfferDates(offerDates []int) {

}
