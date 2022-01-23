package model

type Review struct {
	Base
	Rating      float32 `gorm:"type:float(3);not null"`
	User        *User
	Course      *Course
	Anonymous   bool   `gorm:"type:boolean;not null"`
	Recommended bool   `gorm:"type:boolean;not null"`
	Pros        string `gorm:"type:text;"`
	Cons        string `gorm:"type:text"`
	Comment     string `gorm:"type:text"`
	CourseID    int64  `gorm:"type:bigint;not null"`
	UserID      int64  `gorm:"type:bigint;not null"`
}
