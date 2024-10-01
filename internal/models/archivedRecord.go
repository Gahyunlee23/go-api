package models

import (
	"time"

	"gorm.io/datatypes"
)

type ArchivedRecord struct {
	ID           uint           `gorm:"primary_key;auto_increment" json:"id"`
	SourceTable  string         `gorm:"not null" json:"source_table"`
	ArchivedData datatypes.JSON `gorm:"type:json;not null" swaggerignore:"true" json:"archived_data"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" swaggerignore:"true" json:"created_at"`
}

func (*ArchivedRecord) TableName() string {
	return "archived_record"
}
