package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FileType struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string         `gorm:"type:varchar(255)" json:"code"`
	Name        string         `gorm:"type:varchar(255)" json:"name"`
	Description string         `gorm:"type:mediumtext" json:"description"`
	Extensions  datatypes.JSON `gorm:"not null" swaggerignore:"true" json:"extensions"`
	Price       *float64       `gorm:"type:double;null" json:"price"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
}

func (*FileType) TableName() string {
	return "file_type"
}
