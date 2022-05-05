package migrate

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pmc_server/model"
)

func InitDatabases() (db *gorm.DB, driver neo4j.Driver, err error) {
	dsn := "host=pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com user=admin1 password=admin123 dbname=postgres port=5432 sslmode=disable"
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

func Majors() error {
	db, driver, err := InitDatabases()
	if err != nil {
		return err
	}

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (m:Major{college_id:2}) RETURN m.name, m.emphasis_required", nil)
		if err != nil {
			return nil, err
		}
		majorMap := make(map[string]bool, 0)
		for records.Next() {
			if name, ok := records.Record().Values[0].(string); ok {
				if requireEmphasis, ok := records.Record().Values[1].(bool); ok {
					majorMap[name] = requireEmphasis
				}
			}
		}
		return majorMap, nil
	})

	for k, v := range result.(map[string]bool) {
		major := model.Major{
			CollegeID:        1,
			Name:             k,
			EmphasisRequired: v,
			IsEmphasis:       false,
			MainMajorID:      -1,
		}
		res := db.Create(&major)
		if res.Error != nil {
			return res.Error
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func Emphasis() error {
	db, driver, err := InitDatabases()
	if err != nil {
		return err
	}

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("match (e:Emphasis{college_id:2})-[:SUB_OF]->(major) return e.name, major.name",
			nil)
		if err != nil {
			return nil, err
		}
		emphasisMap := make(map[string]string, 0)
		for records.Next() {
			if name, ok := records.Record().Values[0].(string); ok {
				if majorName, ok := records.Record().Values[1].(string); ok {
					emphasisMap[name] = majorName
				}
			}
		}
		return emphasisMap, nil
	})

	for k, v := range result.(map[string]string) {
		var mainMajor model.Major
		res := db.Where("name = ?", v).First(&mainMajor)
		if res.Error != nil {
			return res.Error
		}
		emphasis := model.Major{
			CollegeID:        1,
			Name:             k,
			EmphasisRequired: false,
			IsEmphasis:       true,
			MainMajorID:      int32(mainMajor.ID),
		}
		res = db.Create(&emphasis)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
