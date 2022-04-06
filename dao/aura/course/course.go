package course

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"pmc_server/init/aura"
)

type Set struct {
	Name           string
	Relation       string
	TargetName     string
	CourseRequired int32
	LinkedToMajor  bool
}

type InsertSet struct {
	Set Set
}

func (s InsertSet) InsertCourseSet() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(s.InsertCourseSetFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (s *InsertSet) InsertCourseSetFn(tx neo4j.Transaction) (interface{}, error) {
	target := ""
	if s.Set.LinkedToMajor {
		target = "Degree"
	} else {
		target = "CourseSet"
	}
	command := fmt.Sprintf("MATCH (m:%s) WHERE m.name = $target_name "+
		"CREATE (s:CourseSet)-[:%s]->(m) SET s.name = $name, s.course_required = $course_required  "+
		"RETURN s.name", target, s.Set.Relation)
	records, err := tx.Run(command,
		map[string]interface{}{
			"target_name":     s.Set.TargetName,
			"relation":        s.Set.Relation,
			"name":            s.Set.Name,
			"course_required": s.Set.CourseRequired,
		})

	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return record.Values[0], nil
}

type Entity struct {
	Name    string
	ID      int64
	SetName string
}

type InsertEntity struct {
	Entity Entity
}

func (s InsertEntity) Insert() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(s.InsertFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (s *InsertEntity) InsertFn(tx neo4j.Transaction) (interface{}, error) {
	command := "MATCH (c:CourseSet) WHERE c.name = $set_name " +
		"CREATE (s:Course)<-[:INCLUDES]-(c) SET s.name = $name, s.id = $id " +
		"RETURN s.name"
	records, err := tx.Run(command, map[string]interface{}{
		"set_name": s.Entity.SetName,
		"name":     s.Entity.Name,
		"id":       s.Entity.ID,
	})

	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return record.Values[0], nil
}
