package models

import "gorm.io/gorm"

type ProductLite struct {
	gorm.Model
	ID         uint   `json:"id"`
	CloudLabID int    `json:"cloud_lab_id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Type       string `json:"type"`
}

func (*ProductLite) TableName() string {
	return "product"
}
