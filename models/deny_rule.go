package models

import (
	"gorm.io/datatypes"
)

type DenyRule struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	ProductPartID *uint          `gorm:"index;foreignKey:ProductPartID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name          string         `gorm:"type:varchar(255);not null"`
	Code          string         `gorm:"type:varchar(255);not null"`
	IsGlobal      bool           `gorm:"type:tinyint(1);not null"`
	Paper         datatypes.JSON `gorm:"type:json;not null"`
	Format        datatypes.JSON `gorm:"type:json;not null"`
	Pages         datatypes.JSON `gorm:"type:json;not null"`
	Colors        datatypes.JSON `gorm:"type:json;not null"`
	BookBinding   datatypes.JSON `gorm:"type:json;not null"`
	Refinement    datatypes.JSON `gorm:"type:json;not null"`
	Finishing     datatypes.JSON `gorm:"type:json;not null"`
	ProductPart   ProductPart    `gorm:"foreignKey:ProductPartID;references:ID"`
}
