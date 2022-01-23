package model

type User struct {
	Base
	UserID       int64   `gorm:"type:bigint(20);unique;not null"`
	FirstName    string  `gorm:"type:varchar(30);not null"`
	LastName     string  `gorm:"type:varchar(30);not null"`
	Password     string  `gorm:"type:varchar(120);not null"`
	Email        string  `gorm:"index:idx_email;type:varchar(30);not null"`
	CollegeID    int64   `gorm:"type:varchar(80)"`
	College      College `gorm:"foreignKey:CollegeID"`
	AcademicYear string  `gorm:"type:varchar(5)"`
	Avatar       string  `gorm:"type:varchar(200)"`
	Role         int32   `gorm:"type:smallint;not null;default 1"`
}
