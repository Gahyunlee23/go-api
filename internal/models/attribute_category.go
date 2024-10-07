package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AttributeCategory struct {
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	Code             string         `gorm:"type:varchar(255);not null;uniqueIndex" json:"code"`
	Name             string         `gorm:"type:varchar(255);not null" json:"name"`
	RequiredSettings datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"required_settings"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Attributes       []Attribute    `gorm:"foreignKey:CategoryID" json:"attributes"`
}

func (*AttributeCategory) TableName() string {
	return "attribute_category"
}
