package major

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"pmc_server/init/aura"
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
