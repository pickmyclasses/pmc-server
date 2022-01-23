package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int64     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}
