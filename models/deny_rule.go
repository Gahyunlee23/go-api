package models

import (
	"gorm.io/datatypes"
)

type DenyRule struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	ProductPartID *uint          `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name          string         `gorm:"type:varchar(255);not null"`
	Code          string         `gorm:"type:varchar(255);not null"`
	IsGlobal      bool           `gorm:"not null"`
	Paper         datatypes.JSON `gorm:"not null"`
	Format        datatypes.JSON `gorm:"not null"`
	Pages         datatypes.JSON `gorm:"not null"`
	Colors        datatypes.JSON `gorm:"not null"`
	BookBinding   datatypes.JSON `gorm:"not null"`
	Refinement    datatypes.JSON `gorm:"not null"`
	Finishing     datatypes.JSON `gorm:"not null"`
	ProductPart   ProductPart    `gorm:"foreignKey:ProductPartID"`
}

func (*DenyRule) TableName() string {
	return "deny_rule"
}
