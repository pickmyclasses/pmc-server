package model

type Department struct {
	Base
	Name         string `gorm:"type:varchar(40);not null;index:idx_name;unique"`
	Abbreviation string `gorm:"type:varchar(10);not null;index:idx_abbr;unique"`
	CollegeID    int64
	College      *College `gorm:"foreignKey:CollegeID"`
}
