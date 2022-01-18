package model

type Course struct {
	Base
	DepartmentID     int64
	PrerequisiteOfID int64
	Department       *Department
	Catalog          string  `gorm:"type:varchar(40);not null"`
	Name             string  `gorm:"type:varchar(50);not null"`
	Number           string  `gorm:"type:varchar(20);not null"`
	Description      string  `gorm:"type:text;not null"`
	SyllabusURL      string  `gorm:"type:varchar(120);"`
	Introduction     string  `gorm:"type:text;not null"`
	Type             string  `gorm:"type:varchar(20);not null"`
	CurrentlyOffered bool    `gorm:"type:boolean;not null"`
	Rating           float32 `gorm:"type:float(3)"`
	PrerequisiteOf   *Course
}
