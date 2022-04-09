package major

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"pmc_server/dao/aura/course"
	"pmc_server/init/aura"
	"pmc_server/shared"
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
	Name          string `json:"name"`
	TotalCredit   int32  `json:"totalCredit"`
	MainMajorName string `json:"mainMajorName"`
	CollegeID     int32  `json:"collegeID"`
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
	records, err := tx.Run("MATCH (m:Emphasis) WHERE m.m = $major_name "+
		"CREATE (e:Emphasis)-[:SUB_OF]->(m) SET e.name = $name, e.total_credit = $total_credit, "+
		"e.college_id = $college_id RETURN e.name",
		map[string]interface{}{
			"major_name":   m.Emphasis.MainMajorName,
			"name":         m.Emphasis.Name,
			"total_credit": m.Emphasis.TotalCredit,
			"college_id":   m.Emphasis.CollegeID,
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

type DegreeType struct {
	Name      string
	Major     string
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
	records, err := tx.Run("MATCH (m:Major) WHERE m.name = $major_name "+
		"CREATE (d:DegreeType)<-[HAS_DEGREE]-(m) SET d.name = $name, d.college_id = $college_id RETURN d.name",
		map[string]interface{}{
			"major_name": m.Type.Major,
			"name":       m.Type.Name,
			"college_id": m.Type.CollegeID,
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

type Read struct {
	CollegeID int32
}

func (r Read) FindAll() ([]Entity, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.ReadTransaction(r.findAllFn)
	if err != nil {
		return nil, err
	}

	return result.([]Entity), nil
}

func (r *Read) findAllFn(tx neo4j.Transaction) (interface{}, error) {
	res, err := tx.Run("MATCH (m:Major {college_id : $college_id})"+
		"RETURN m.degree_hour, m.emphasis_required, m.min_major_hour, m.name",
		map[string]interface{}{
			"college_id": r.CollegeID,
		})

	if err != nil {
		return nil, err
	}

	majorList := make([]Entity, 0)
	for res.Next() {
		var major Entity
		major.CollegeID = int(r.CollegeID)
		if name, ok := res.Record().Values[3].(string); ok {
			major.Name = name
		}
		if degreeHour, ok := res.Record().Values[0].(int64); ok {
			major.DegreeHour = int32(degreeHour)
		}
		if minMajorHour, ok := res.Record().Values[2].(int64); ok {
			major.MinMajorHour = int32(minMajorHour)
		}
		if emphasisRequired, ok := res.Record().Values[1].(bool); ok {
			major.EmphasisRequired = emphasisRequired
		}

		majorList = append(majorList, major)
	}

	return majorList, nil
}

type ReadEmphasis struct {
	CollegeID int32
	MajorName string
}

func (r ReadEmphasis) FindAllEmphasisesOfAMajor() ([]Emphasis, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.ReadTransaction(r.findAllEmphasisesFn)
	if err != nil {
		return nil, err
	}

	return result.([]Emphasis), nil
}

func (r *ReadEmphasis) findAllEmphasisesFn(tx neo4j.Transaction) (interface{}, error) {
	res, err := tx.Run("MATCH (m:Major {college_id: $college_id, name: $major_name})<-[:SUB_OF]-(emphasis) "+
		" RETURN emphasis.name, emphasis.total_credit",
		map[string]interface{}{
			"college_id": r.CollegeID,
			"major_name": r.MajorName,
		})

	if err != nil {
		return nil, err
	}

	emphasisList := make([]Emphasis, 0)
	for res.Next() {
		var emphasis Emphasis
		if name, ok := res.Record().Values[0].(string); ok {
			emphasis.Name = name
		}
		if totalCredit, ok := res.Record().Values[1].(int32); ok {
			emphasis.TotalCredit = totalCredit
		}
		emphasis.CollegeID = r.CollegeID
		emphasis.MainMajorName = r.MajorName
		emphasisList = append(emphasisList, emphasis)
	}

	return emphasisList, nil
}

// Reader is for reading the course entity list of a course set
// This will give the entire list of the course list under a course set
type Reader struct {
	MajorName string // the name of major we want to fetch
	SetName   string // the name of the set we are fetching (this could be empty)
}

// ReadList defines a reader for reading the course list
type ReadList struct {
	Reader Reader
}

// ReadAll reads the course list from a course set
func (r ReadList) ReadAll() ([]int64, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(r.ReadAllFn)
	if err != nil {
		return nil, err
	}
	return result.([]int64), nil
}

// ReadAllFn is a helper function of ReadAll
func (r *ReadList) ReadAllFn(tx neo4j.Transaction) (interface{}, error) {
	res, err := tx.Run("MATCH (m:Major)-[*]-(connected) RETURN connected", map[string]interface{}{
		"major_name": r.Reader.MajorName,
	})
	if err != nil {
		return nil, err
	}

	courseList := make([]int64, 0)
	for res.Next() {
		fmt.Println(res.Record().Values[0])
	}

	return courseList, nil
}

func (r ReadList) ReadDirectCourseSet() ([]course.Set, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(r.ReadDirectCourseSetFn)
	if err != nil {
		return nil, err
	}
	return result.([]course.Set), nil
}

func (r ReadList) ReadDirectCourseSetFn(tx neo4j.Transaction) (interface{}, error) {
	command := fmt.Sprintf("MATCH (m:Major{name:'%s'})-[:HAS]->(d:Degree{name:'Bachelor or Arts - %s'})"+
		"<-[:REQUIRED_BY]-(courseSet) RETURN courseSet.name, courseSet.course_required",
		r.Reader.MajorName, r.Reader.MajorName)
	res, err := tx.Run(command, nil)
	if err != nil {
		return nil, err
	}

	courseSetList := make([]course.Set, 0)
	for res.Next() {
		fmt.Println(res.Record().Values)
	}
	return courseSetList, nil
}

type SubSet struct {
	Name           string  `json:"name"`
	CourseRequired int32   `json:"courseRequired"`
	CourseIDList   []int64 `json:"courseIDList"`
}

func (r ReadList) ReadSubCourseSets() ([]SubSet, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	result, err := session.ReadTransaction(r.ReadSubCourseSetsFn)
	if err != nil {
		return nil, err
	}
	return result.([]SubSet), nil
}

func (r ReadList) ReadSubCourseSetsFn(tx neo4j.Transaction) (interface{}, error) {
	command := fmt.Sprintf(
		"MATCH (connected)-[*]-(m:CourseSet{name:\"%s\"})-[:REQUIRED_BY]->(:Degree{name:\"Bachelor or Arts -"+
			" Accounting\"})<-[:HAS]-(major:Major{name:\"%s\"}) "+
			"RETURN connected.name, labels(connected), connected.course_required, connected.id",
		r.Reader.SetName, r.Reader.MajorName)
	res, err := tx.Run(command, nil)
	if err != nil {
		return nil, err
	}

	subSetList := make([]SubSet, 0)
	isFirstSet := true
	isPrevSet := false
	for res.Next() {
		var subSet SubSet
		if labels, ok := res.Record().Values[1].([]interface{}); ok {
			if !isFirstSet {
				subSetList = append(subSetList, subSet)
			}
			// handle courseSet
			if labels[0].(string) == "CourseSet" {
				// previous one is also a set, this is a subset
				if isPrevSet {
					isPrevSet = true
				}
				subSet = SubSet{}
				subSet.Name = res.Record().Values[0].(string)
				fmt.Println(res.Record().Values[2])
				subSet.CourseRequired = int32(res.Record().Values[2].(int64))
				subSet.CourseIDList = make([]int64, 0)
				isFirstSet = false
			}
			if labels[0].(string) == "Course" {
				isPrevSet = false
				subSet.CourseIDList = append(subSet.CourseIDList, res.Record().Values[3].(int64))
			}
		} else {
			return nil, shared.InternalErr{}
		}
	}
	return subSetList, nil
}
