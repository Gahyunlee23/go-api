package models

import "gorm.io/gorm"

type ProductLite struct {
	ID         uint           `json:"id"`
	CloudLabID int            `json:"cloud_lab_id"`
	Name       string         `json:"name"`
	Code       string         `json:"code"`
	Type       string         `json:"type"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func (*ProductLite) TableName() string {
	return "product"
}
