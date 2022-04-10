package model

import "github.com/lib/pq"

type CourseSet struct {
	Base
	Name           string        `gorm:"name"`
	IsLeaf         bool          `gorm:"is_leaf"`
	CourseIDList   pq.Int64Array `gorm:"column:course_id_list;type:integer[]"`
	ParentSetID    int32         `gorm:"parent_set_id"`
	MajorID        int32         `gorm:"major_id"`
	LinkedToMajor  bool          `gorm:"linked_to_major"`
	CourseRequired int32         `gorm:"course_required"`
}
