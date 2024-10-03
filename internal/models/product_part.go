package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProductPart struct {
	ID                uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name              string         `gorm:"size:255;not null" json:"name"`
	Code              string         `gorm:"size:255;not null;unique" json:"code"`
	ContentType       string         `gorm:"size:255;not null" json:"content_type"`
	Paper             datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"paper"`
	Format            datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"format"`
	Pages             datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"pages"`
	Colors            datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"colors"`
	BookBinding       datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"book_binding"`
	Refinement        datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"refinement"`
	Finishing         datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"finishing"`
	DefaultSelections datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"default_selections"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	DenyRules         []DenyRule     `gorm:"foreignKey:ProductPartID" json:"deny_rules"`
}

func (*ProductPart) TableName() string {
	return "product_part"
}
