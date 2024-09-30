package models

import (
	"time"

	"gorm.io/datatypes"
)

type ArchivedRecord struct {
	ID           uint           `gorm:"primary_key;auto_increment"`
	SourceTable  string         `gorm:"not null"`
	ArchivedData datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" swaggerignore:"true"`
}

func (*ArchivedRecord) TableName() string {
	return "archived_record"
}
