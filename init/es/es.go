package es

import (
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

var Client *elastic.Client

func Init(url, username, password string) error {
	var err error
	Client, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetBasicAuth(username, password))
	if err != nil {
		zap.L().Error("init elastic search failed")
		return err
	}
	return nil
}
