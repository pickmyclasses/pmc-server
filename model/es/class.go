package es

type Class struct {
	ID         int64   `json:"id"`
	CourseID   int64   `json:"course_id"`
	IsOnline   bool    `json:"is_online"`
	IsInPerson bool    `json:"is_in_person"`
	IsHybrid   bool    `json:"is_hybrid"`
	IsIVC      bool    `json:"is_ivc"`
	OfferDates []int   `json:"offer_dates"`
	StartTime  float32 `json:"start_time"`
	EndTime    float32 `json:"end_time"`
	Professors string  `json:"professors"`
}

func (Class) GetMapping() string {
	return `
{
   "mappings":{
         "properties":{
            "id":{
               "type":"integer"
            },
            "course_id":{
               "type":"integer"
            },
            "is_online":{
               "type":"boolean"
            },
            "is_in_person":{
               "type":"boolean"
            },
            "is_hybrid":{
               "type":"boolean"
            },
			"is_ivc":{
				"type":"boolean"
			},
            "offer_dates":{
               "type":"nested",
 			   "properties":{  
               		"date_num":{ "type":"integer"}
            	}
            },
            "start_time":{
               "type":"float"
            },
            "end_time":{
               "type":"float"
            },
            "professors":{
               "type":"text"
            }
         }
      }
}
`
}
