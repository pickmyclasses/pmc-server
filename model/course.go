package model

type Course struct {
	Base
	IsHonor            bool     `gorm:"column:is_honor" json:"isHonor"`
	FixedCredit        bool     `gorm:"column:fixed_credit" json:"fixedCredit"`
	DesignationCatalog string   `gorm:"column:designation_catalog" json:"designationCatalog"`
	Description        string   `gorm:"column:description" json:"description"`
	Prerequisites      string   `gorm:"column:prerequisites" json:"prerequisites"`
	Title              string   `gorm:"column:title" json:"title"`
	CatalogCourseName  string   `gorm:"column:catalog_course_name" json:"catalogCourseName"`
	Component          string   `gorm:"column:component" json:"component"`
	MaxCredit          string   `gorm:"column:max_credit" json:"maxCredit"`
	MinCredit          string   `gorm:"column:min_credit" json:"minCredit"`
	SubjectID          int64    `gorm:"column:subject_id" json:"subjectID"`
	AssociatedCourseID int64    `gorm:"column:associated_course_id" json:"associatedCourseID"`
	Subject            *Subject `gorm:"foreignKey:SubjectID" json:"subject"`
}
