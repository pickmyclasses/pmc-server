package model

import (
	"time"
)

type Class struct {
	Base
	Course              *Course
	CourseID            int64      `gorm:"type:bigint,not null"`
	Semester            string     `gorm:"type:varchar(10);not null"`
	Year                string     `gorm:"type:varchar(5);not null"`
	Section             string     `gorm:"type:varchar(10);not null"`
	WaitList            bool       `gorm:"type:boolean;not null"`
	OfferDate           string     `gorm:"type:varchar(30);not nul"`
	StartTime           *time.Time `gorm:"not null"`
	EndTime             *time.Time `gorm:"not null"`
	Location            string     `gorm:"type:varchar(30);not null"`
	RecommendationScore float32    `gorm:"not null"`
	Type                int8       `gorm:"type:smallint;not null"`
}
