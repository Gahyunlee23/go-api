package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AttributeCategory struct {
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	Code             string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	Name             string         `gorm:"type:varchar(255);not null"`
	RequiredSettings datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

func (*AttributeCategory) TableName() string {
	return "attribute_category"
}
