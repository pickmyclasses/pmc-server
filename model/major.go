package model

type Major struct {
	Base
	CollegeID        int32  `gorm:"college_id"`
	Name             string `gorm:"column:name"`
	EmphasisRequired bool   `gorm:"column:emphasis_required"`
	IsEmphasis       bool   `gorm:"is_emphasis"`
	MainMajorID      int32  `gorm:"main_major_id"`
}
