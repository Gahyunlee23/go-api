package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	ID                   uint           `gorm:"primaryKey;autoIncrement"`
	CloudLabID           int            `json:"cloud_lab_id" gorm:"not null"`
	Name                 string         `json:"name" gorm:"size:255;not null"`
	Code                 string         `json:"code" gorm:"size:255;not null;unique"`
	Type                 string         `json:"type" gorm:"size:255;not null"`
	Description          string         `json:"description" gorm:"type:mediumtext"`
	MinimumQuantity      int            `json:"minimum_quantity" gorm:"default:1;not null"`
	MaximumQuantity      *int           `json:"maximum_quantity"`
	PackingUnit          int            `json:"packing_unit" gorm:"default:1;not null"`
	EnableCustomQuantity bool           `json:"enable_custom_quantity" gorm:"default:0;not null"`
	EnableCustomFormat   bool           `json:"enable_custom_format" gorm:"default:0;not null"`
	TimeToProduce        *string        `json:"time_to_produce" gorm:"size:255"`
	RenamingRules        datatypes.JSON `json:"renaming_rules" swaggerignore:"true" gorm:"type:json;not null;"`
	OrderRules           datatypes.JSON `json:"order_rules" swaggerignore:"true" gorm:"type:json;not null"`
	DefaultQuantity      int            `json:"default_quantity" gorm:"not null"`
	QuantitiesSelection  datatypes.JSON `json:"quantities_selection" swaggerignore:"true" gorm:"type:json;not null"`
	PriceCalculationType string         `json:"price_calculation_type" gorm:"size:50;not null"`
	CreatedAt            time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt            gorm.DeletedAt `json:"deleted_at" swaggerignore:"true" gorm:"index"`
}

func (*Product) TableName() string {
	return "product"
}
