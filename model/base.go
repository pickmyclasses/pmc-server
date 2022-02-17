package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	IsDeleted bool           `gorm:"column:is_deleted" json:"is_deleted"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
