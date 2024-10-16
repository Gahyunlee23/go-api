package models

import (
	"time"

	"gorm.io/gorm"
)

type Proof struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string         `gorm:"type:varchar(255)" json:"code"`
	Name        string         `gorm:"type:varchar(255)" json:"name"`
	Description string         `gorm:"type:mediumtext" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"deleted_at"`
	Product     []Product      `gorm:"many2many:product_proof"`
}

func (*Proof) TableName() string {
	return "proof"
}
