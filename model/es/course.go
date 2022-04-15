package es

type Course struct {
	ID                 int64         `json:"id"`
	DesignationCatalog string        `json:"designation_catalog"`
	Title              string        `json:"title"`
	Description        string        `json:"description"`
	CatalogCourseName  string        `json:"catalog_course_name"`
	CatalogName        string        `json:"catalog_name"`
	Prerequisites      string        `json:"prerequisites"`
	Component          string        `json:"component"`
	MaxCredit          float32       `json:"max_credit"`
	MinCredit          float32       `json:"min_credit"`
	Subject            string        `json:"subject"`
	IsHonor            bool          `json:"is_honor"`
	FixedCredit        bool          `json:"fixed_credit"`
	Rating             float32       `json:"rating"`
	HasOffering        bool          `json:"hasOffering"`
	Classes            []ClassNested `json:"classes" elastic:"type:nested"`
	Tags               []string      `json:"tags"`
}

type ClassNested struct {
	Semester    string  `json:"semester"`
	Year        int     `json:"year"`
	WaitList    string  `json:"wait_list"`
	OfferDate   string  `json:"offer_date"`
	StartTime   float32 `json:"start_time"`
	EndTime     float32 `json:"end_time"`
	Instructors string  `json:"instructors"`
}

func (Course) GetMapping() string {
	return `
	{
  "mappings":{
		"course": {
        "properties":{
           "id":{
              "type":"integer"
           },
           "title":{
              "type":"text"
           },
           "description":{
              "type":"text"
           },
           "designation_catalog":{
              "type":"text"
           },
           "catalog_course_name":{
              "type":"text"
           },
			      "catalog_name":{
				      "type":"text"
			      },
           "prerequisites":{
              "type":"text"
           },
           "component":{
              "type":"text"
           },
           "max_credit":{
              "type":"integer"
           },
           "min_credit":{
              "type":"integer"
           },
           "subject":{
              "type":"text"
           },
           "is_honor":{
              "type":"boolean"
           },
           "fixed_credit":{
              "type":"boolean"
           },
			      "rating": {
				      "type":"float"
			      },
			      "has_offering":{
				      "type":"boolean"
			      },
			      "classes": {
				      "type":"nested"
			      },
			      "tags": {
				      "type": "text"
			      }
        }
     }
	}
}


`
}

func (Course) GetIndexName() string {
	return "course"
}
