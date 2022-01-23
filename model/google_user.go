package model

import "time"

type GoogleUser struct {
	Base
	UserID        int64     `gorm:"type:bigint;not null"`
	FirstName     string    `gorm:"type:varchar(30);not null"`
	LastName      string    `gorm:"type:varchar(30);not null"`
	Email         string    `gorm:"type:varchar(40);not null"`
	Avatar        string    `gorm:"type:varchar(30);not null"`
	AccessToken   string    `gorm:"type:varchar(100);not null"`
	TokenExpireAt time.Time `gorm:"type:timestamp;not null"`
	RefreshToken  string    `gorm:"type:varchar(3100);not null"`
}
