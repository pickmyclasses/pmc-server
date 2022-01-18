package model

type College struct {
	Base
	Name string `gorm:"type:char(30);not null"`
}
