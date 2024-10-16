package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductionTime struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Code      string         `gorm:"type:varchar(255);not null" json:"code"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Time      string         `gorm:"type:varchar(255);not null" json:"time"`
	Price     *float64       `gorm:"type:double;null" json:"price"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Product   []Product      `gorm:"many2many:product_proof"`
}

func (*ProductionTime) TableName() string {
	return "production_time"
}
