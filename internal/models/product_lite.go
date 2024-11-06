package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductLite struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CloudLabID int            `gorm:"not null" json:"cloud_lab_id"`
	Name       string         `gorm:"size:255;not null" json:"name"`
	Code       string         `gorm:"size:255;not null;unique" json:"code"`
	Type       string         `gorm:"size:255;not null" json:"type"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

func (*ProductLite) TableName() string {
	return "product"
}
