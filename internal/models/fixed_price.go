package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FixedPrice struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductPartID *uint          `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product_part_id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Code          string         `gorm:"type:varchar(255);not null" json:"code"`
	Prices        datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"prices"`
	Paper         datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"paper"`
	Format        datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"format"`
	Pages         datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"pages"`
	Colors        datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"colors"`
	BookBinding   datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"book_binding"`
	Refinement    datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"refinement"`
	Finishing     datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"finishing"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	ProductPart   ProductPart    `gorm:"foreignKey:product_part_id"`
}

func (*FixedPrice) TableName() string {
	return "fixed_price"
}
