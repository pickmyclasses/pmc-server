package aura

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

var Driver neo4j.Driver

func Init(uri, username, password string) error {
	var err error
	auth := neo4j.BasicAuth(username, password, "")
	Driver, err = neo4j.NewDriver(uri, auth)
	if err != nil {
		return err
	}
	return nil
}
