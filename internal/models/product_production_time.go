package models

type ProductProductionTime struct {
	ProductID uint `gorm:"primaryKey"`
	ProofID   uint `gorm:"primaryKey"`
	IsDefault bool `gorm:"default:0;not null" json:"is_default"`
}

func (*ProductProductionTime) TableName() string {
	return "product_production_time"
}
