package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                   uint           `gorm:"primaryKey;autoIncrement"`
	CloudLabID           int            `gorm:"not null"`
	Name                 string         `gorm:"size:255;not null"`
	Code                 string         `gorm:"size:255;not null;unique"`
	Type                 string         `gorm:"size:255;not null"`
	Description          string         `gorm:"type:mediumtext"`
	MinimumQuantity      int            `gorm:"default:1;not null"`
	MaximumQuantity      *int           `gorm:""`
	PackingUnit          int            `gorm:"default:1;not null"`
	EnableCustomQuantity bool           `gorm:"default:0;not null"`
	EnableCustomFormat   bool           `gorm:"default:0;not null"`
	TimeToProduce        *string        `gorm:"size:255"`
	RenamingRules        datatypes.JSON `gorm:"type:json;not null"`
	OrderRules           datatypes.JSON `gorm:"type:json;not null"`
	DefaultQuantity      int            `gorm:"not null"`
	QuantitiesSelection  datatypes.JSON `gorm:"type:json;not null"`
	PriceCalculationType string         `gorm:"size:50;not null"`
}
