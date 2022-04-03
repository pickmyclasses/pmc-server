package major

import (
	"fmt"
	"pmc_server/init/aura"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Entity struct {
	CollegeID        int
	Name             string
	DegreeHour       int32
	MinMajorHour     int32
	EmphasisRequired bool
}

type Insertion struct {
	Major Entity
}

func (m Insertion) InsertMajor() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(m.insertMajorFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (m *Insertion) insertMajorFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("CREATE (m:Major { name: $name, degree_hour: $degree_hour,"+
		" min_major_hour: $min_major_hour, emphasis_required: $emphasis_required, college_id: $college_id}) "+
		"RETURN m.college_id, m.name, m.degree_hour, m.min_major_hour, m.emphasis_required",
		map[string]interface{}{
			"college_id":        m.Major.CollegeID,
			"name":              m.Major.Name,
			"degree_hour":       m.Major.DegreeHour,
			"min_major_hour":    m.Major.MinMajorHour,
			"emphasis_required": m.Major.EmphasisRequired,
		})

	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return record.Values[1], nil
}

type Emphasis struct {
	Name string
	TotalCredit int32
	MainMajorName string
	CollegeID int32
}

type EmphasisInsertion struct {
	Emphasis Emphasis
}

func (m EmphasisInsertion) InsertEmphasis() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(m.insertEmphasisFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (m *EmphasisInsertion) insertEmphasisFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("MATCH (m:Major) WHERE m.name = $major_name " +
		"CREATE (e:Emphasis)-[:SUB_OF]->(m) SET e.name = $name, e.total_credit = $total_credit, " +
		"e.college_id = $college_id RETURN e.name",
		map[string]interface{}{
		"major_name": m.Emphasis.MainMajorName,
		"name": m.Emphasis.Name,
		"total_credit": m.Emphasis.TotalCredit,
		"college_id": m.Emphasis.CollegeID,
		})

	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	fmt.Println(record)
	if err != nil {
		return nil, err
	}

	return record.Values[0], nil
}

type DegreeType struct {
	Name string
	Major string
	CollegeID int32
}

type DegreeInsertion struct {
	Type DegreeType
}

func (m DegreeInsertion) InsertDegreeType() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(m.insertDegreeTypeFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (m *DegreeInsertion) insertDegreeTypeFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("MATCH (m:Major) WHERE m.name = $major_name " +
		"CREATE (d:DegreeType)<-[HAS_DEGREE]-(m) SET d.name = $name, d.college_id = $college_id RETURN d.name",
		map[string]interface{}{
			"major_name": m.Type.Major,
			"name": m.Type.Name,
			"college_id": m.Type.CollegeID,
		})
	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	fmt.Println(record)
	if err != nil {
		return nil, err
	}

	return record.Values[0], nil
}


