package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Attribute struct {
	ID          uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID  uint              `gorm:"not null;index" json:"category_id"`
	Code        string            `gorm:"type:varchar(255);not null;uniqueIndex" json:"code"`
	Name        string            `gorm:"type:varchar(255);not null" json:"name"`
	Description string            `gorm:"type:mediumtext" json:"description"`
	Order       int               `gorm:"not null" json:"order"`
	Settings    datatypes.JSON    `gorm:"type:json" swaggerignore:"true" json:"settings"`
	CreatedAt   time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Category    AttributeCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (*Attribute) TableName() string {
	return "attribute"
}
