package migrate

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDatabases() (db *gorm.DB, driver neo4j.Driver, err error) {
	dsn := "host=localhost user=postgres password=admin123 dbname=pmc port=5432 sslmode=disable"
	db, err = gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	auth := neo4j.BasicAuth("neo4j", "9_a9iCvVKRvW8LbdT5eiBfeSFQdclfboXk0JcuZ9mgQ", "")
	driver, err = neo4j.NewDriver("neo4j+s://f8e9dc6c.databases.neo4j.io:7687", auth)
	if err != nil {
		return
	}
	return
}
