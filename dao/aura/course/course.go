// Package course - aura dao for course
// All rights reserved by pickmyclass.com
// Author: Kaijie Fu
// Date: 3/13/2022
package course

import (
	"fmt"

	"pmc_server/init/aura"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Set defines a course SET, for example, CS pre-major
// Set contains multiple courses
// This is used for inserting mostly
type Set struct {
	Name           string // name of the course set
	Relation       string // relationship will be attached to another node
	TargetName     string // the target node this set will be connected to, cannot be null
	CourseRequired int32  // how many courses are required in the set (minimum credits required)
	LinkedToMajor  bool   // is this set directly linked to the major node (sometimes they are subset of another set)
}

// InsertSet defines the action of inserting a course set
type InsertSet struct {
	Set Set
}

// InsertCourseSet defines the behavior of inserting a course set to Neo4j
func (s InsertSet) InsertCourseSet() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(s.InsertCourseSetFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

// InsertCourseSetFn is a helper function for InsertCourseSet
func (s *InsertSet) InsertCourseSetFn(tx neo4j.Transaction) (interface{}, error) {
	target := ""
	// check if it's linked to a major directly
	// here I'm calling it Degree because a course set needs to be linked to one of the degrees of the major
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

// Entity defines a course entity node
type Entity struct {
	Name    string // the name of the course
	ID      int64  // the ID of the course, this will be used for fetching the actual course entity from Postgres
	SetName string
}

// InsertEntity defines the action to insert a course entity node to Neo4j
type InsertEntity struct {
	Entity Entity
}

// Insert defines the behavior of inserting a course node to Neo4j
func (s InsertEntity) Insert() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(s.InsertFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

// InsertFn is the helper function for Insert
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
