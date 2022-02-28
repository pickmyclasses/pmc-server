package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	IsDeleted bool      `gorm:"column:is_deleted" json:"isDeleted"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	DeletedAt gorm.DeletedAt
}
