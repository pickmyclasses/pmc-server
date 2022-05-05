package model

import "github.com/lib/pq"

type CustomEvent struct {
	Base
	Title       string        `gorm:"column:title"`
	Description string        `gorm:"column:description"`
	Color       string        `gorm:"column:color"`
	Days        pq.Int64Array `gorm:"column:days;type:integer[]"`
	StartTime   int32         `gorm:"column:start_time"`
	EndTime     int32         `gorm:"column:end_time"`
	UserID      int64         `gorm:"column:user_id"`
	SemesterID  int64         `gorm:"column:semester_id"`
	Kind        string        `gorm:"kind"`
}
