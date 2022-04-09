package model

type User struct {
	Base
	UserID       int64   `gorm:"type:bigint(20);unique;not null"` // snowflake ID
	FirstName    string  `gorm:"type:varchar(30);not null"`
	LastName     string  `gorm:"type:varchar(30);not null"`
	Password     string  `gorm:"type:varchar(120);not null"`
	Email        string  `gorm:"index:idx_email;type:varchar(30);not null"`
	CollegeID    int64   `gorm:"type:varchar(80)"`
	College      College `gorm:"foreignKey:CollegeID"`
	AcademicYear string  `gorm:"type:varchar(5)"`                  // user's school year, 1-4, default 1
	Avatar       string  `gorm:"type:varchar(200)"`                // user's avatar address, on S3
	Role         int32   `gorm:"type:smallint;not null;default 1"` // user's role, 0 admin, 1 student, 2 professor
	Major        string  `gorm:"column:major"`                     // student major, this could be empty if not provided by the user
	Emphasis     string  `gorm:"column:emphasis"`                  // student emphasis (if any)
}
