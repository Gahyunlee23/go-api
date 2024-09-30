package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DenyRule struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	ProductPartID *uint          `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name          string         `gorm:"type:varchar(255);not null"`
	Code          string         `gorm:"type:varchar(255);not null"`
	IsGlobal      bool           `gorm:"not null"`
	Paper         datatypes.JSON `gorm:"not null" swaggerignore:"true"`
	Format        datatypes.JSON `gorm:"not null" swaggerignore:"true"`
	Pages         datatypes.JSON `gorm:"not null" swaggerignore:"true"`
	Colors        datatypes.JSON `gorm:"not null" swaggerignore:"true"`
	BookBinding   datatypes.JSON `gorm:"not null" swaggerignore:"true"`
	Refinement    datatypes.JSON `gorm:"not null" swaggerignore:"true"`
	Finishing     datatypes.JSON `gorm:"not null" swaggerignore:"true"`
	ProductPart   ProductPart    `gorm:"foreignKey:ProductPartID"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

func (*DenyRule) TableName() string {
	return "deny_rule"
}
