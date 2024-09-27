package models

import (
	"time"

	"gorm.io/gorm"
)

type Attribute struct {
	ID          uint              `gorm:"primaryKey;autoIncrement"`
	CategoryID  uint              `gorm:"not null;index"`
	Code        string            `gorm:"type:varchar(255);not null;uniqueIndex"`
	Name        string            `gorm:"type:varchar(255);not null"`
	Description string            `gorm:"type:mediumtext"`
	Order       int               `gorm:"not null"`
	CreatedAt   time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Category    AttributeCategory `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (*Attribute) TableName() string {
	return "attribute"
}
