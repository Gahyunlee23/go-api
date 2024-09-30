package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DenyRule struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductPartID *uint          `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Code          string         `gorm:"type:varchar(255);not null" json:"code"`
	IsGlobal      bool           `gorm:"not null" json:"is_global"`
	Paper         datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"paper"`
	Format        datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"format"`
	Pages         datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"pages"`
	Colors        datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"colors"`
	BookBinding   datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"book_binding"`
	Refinement    datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"refinement"`
	Finishing     datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"finishing"`
	ProductPart   ProductPart    `gorm:"foreignKey:ProductPartID" json:"product_part"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

func (*DenyRule) TableName() string {
	return "deny_rule"
}
