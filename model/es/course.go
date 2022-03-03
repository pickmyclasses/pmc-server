package es

type Course struct {
	ID                 int64   `json:"id"`
	DesignationCatalog string  `json:"designation_catalog"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	CatalogCourseName  string  `json:"catalog_course_name"`
	Prerequisites      string  `json:"prerequisites"`
	Component          string  `json:"component"`
	MaxCredit          float32 `json:"max_credit"`
	MinCredit          float32 `json:"min_credit"`
	Subject            string  `json:"subject"`
	IsHonor            bool    `json:"is_honor"`
	FixedCredit        bool    `json:"fixed_credit"`
}

func (Course) GetMapping() string {
	return `
{
   "mappings":{
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
            "prerequisites":{
               "type":"text"
            },
            "component":{
               "type":"text"
            },
            "max_credit":{
               "type":"number"
            },
            "min_credit":{
               "type":"number"
            },
            "subject":{
               "type":"text"
            },
            "is_honor":{
               "type":"boolean"
            },
            "fixed_credit":{
               "type":"boolean"
            }
         }
      }
}
`
}

func (Course) GetIndexName() string {
	return "course"
}
