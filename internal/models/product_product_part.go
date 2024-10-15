package models

type ProductProductPart struct {
	ProductID     uint `gorm:"primaryKey"`
	ProductPartID uint `gorm:"primaryKey"`
}

func (ProductProductPart) TableName() string {
	return "product_product_part"
}
