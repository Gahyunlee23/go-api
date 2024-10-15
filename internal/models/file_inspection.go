package models

import (
	"time"

	"gorm.io/gorm"
)

type FileInspection struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(255)" json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

func (*FileInspection) TableName() string {
	return "file_inspection"
}
