package model

type Course struct {
	Base
	IsHonor            bool     `gorm:"column:is_honor"`
	FixedCredit        bool     `gorm:"column:fixed_credit"`
	DesignationCatalog string   `gorm:"column:designation_catalog"`
	Description        string   `gorm:"column:description"`
	Prerequisites      string   `gorm:"column:prerequisites"`
	Title              string   `gorm:"column:title"`
	CatalogCourseName  string   `gorm:"column:catalog_course_name"`
	Component          string   `gorm:"column:component"`
	MaxCredit          string   `gorm:"column:max_credit"`
	MinCredit          string   `gorm:"column:min_credit"`
	SubjectID          int64    `gorm:"column:subject_id"`
	AssociatedCourseID int64    `gorm:"column:associated_course_id"`
	Subject            *Subject `gorm:"foreignKey:SubjectID"`
}
