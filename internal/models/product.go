package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	ID                   uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CloudLabID           int            `gorm:"not null" json:"cloud_lab_id"`
	Name                 string         `gorm:"size:255;not null" json:"name"`
	Code                 string         `gorm:"size:255;not null;unique" json:"code"`
	Type                 string         `gorm:"size:255;not null" json:"type"`
	Description          string         `gorm:"type:mediumtext" json:"description"`
	MinimumQuantity      int            `gorm:"default:1;not null" json:"minimum_quantity"`
	MaximumQuantity      *int           `gorm:"" json:"maximum_quantity"`
	PackingUnit          int            `gorm:"default:1;not null" json:"packing_unit"`
	EnableCustomQuantity bool           `gorm:"default:0;not null" json:"enable_custom_quantity"`
	EnableCustomFormat   bool           `gorm:"default:0;not null" json:"enable_custom_format"`
	TimeToProduce        *string        `gorm:"size:255" json:"time_to_produce"`
	RenamingRules        datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"renaming_rules"`
	OrderRules           datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"order_rules"`
	DefaultQuantity      int            `gorm:"not null" json:"default_quantity"`
	QuantitiesSelection  datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"quantities_selection"`
	PriceCalculationType string         `gorm:"size:50;not null" json:"price_calculation_type"`
	CreatedAt            time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Parts                []ProductPart  `gorm:"many2many:product_product_part;"`
}

func (*Product) TableName() string {
	return "product"
}
