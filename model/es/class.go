package es

import (
	"time"
)

type Class struct {
	ID         int64     `json:"id"`
	CourseID   int64     `json:"course_id"`
	IsOnline   bool      `json:"is_online"`
	IsInPerson bool      `json:"is_in_person"`
	IsHybrid   bool      `json:"is_hybrid"`
	Rating     float32     `json:"rating"`
	OfferDates []int  	`json:"offer_dates"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Professors []string  `json:"professors"`
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
               "type":"text"
            },
            "is_hybrid":{
               "type":"text"
            },
            "rating":{
               "type":"number"
            },
            "offer_dates":{
               "type":"nested",
				"properties":{  
               		"date_num":{ "type":"integer"},
            	}
            },
            "start_time":{
               "type":"number"
            },
            "end_time":{
               "type":"number"
            },
            "professors":{
               "type":"text"
            },
         }
      }
}
`
}