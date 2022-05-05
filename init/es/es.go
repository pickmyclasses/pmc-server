package es

import (
	"context"
	"errors"
	"log"
	"os"

	model "pmc_server/model/es"

	"github.com/olivere/elastic/v7"
)

var Client *elastic.Client

func Init(url, username, password string) error {
	var err error
	logger := log.New(os.Stdout, "pmc", log.LstdFlags)
	Client, err = elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		//elastic.SetBasicAuth(username, password),
		elastic.SetTraceLog(logger),
	)
	if err != nil {
		return err
	}

	courseIndex := model.Course{}
	exists, err := Client.IndexExists(courseIndex.GetIndexName()).Do(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		if err := CreateIndex(courseIndex.GetIndexName(), courseIndex.GetMapping()); err != nil {
			return err
		}
	}

	return nil
}

func CreateIndex(index, mapping string) error {
	createIndex, err := Client.CreateIndex(index).BodyString(mapping).Do(context.Background())
	if err != nil {
		return err
	}
	if !createIndex.Acknowledged {
		return errors.New("create index failed")
	}
	return nil
}
