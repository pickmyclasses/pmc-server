package model

type Course struct {
	Base
	IsHonor            bool     `gorm:"column:is_honor" json:"is_honor"`
	FixedCredit        bool     `gorm:"column:fixed_credit" json:"fixed_credit"`
	DesignationCatalog string   `gorm:"column:designation_catalog" json:"designation_catalog"`
	Description        string   `gorm:"column:description" json:"description"`
	Prerequisites      string   `gorm:"column:prerequisites" json:"prerequisites"`
	Title              string   `gorm:"column:title" json:"title"`
	CatalogCourseName  string   `gorm:"column:catalog_course_name" json:"catalog_course_name"`
	Component          string   `gorm:"column:component" json:"component"`
	MaxCredit          string   `gorm:"column:max_credit" json:"max_credit"`
	MinCredit          string   `gorm:"column:min_credit" json:"min_credit"`
	SubjectID          int64    `gorm:"column:subject_id" json:"subject_id"`
	AssociatedCourseID int64    `gorm:"column:associated_course_id" json:"associated_course_id"`
	Subject            *Subject `gorm:"foreignKey:SubjectID" json:"subject"`
}
