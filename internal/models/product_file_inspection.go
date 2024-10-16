package models

type ProductFileInspection struct {
	ProductID        uint `gorm:"primaryKey"`
	FileInspectionID uint `gorm:"primaryKey"`
	IsDefault        bool `gorm:"default:0;not null" json:"is_default"`
}

func (*ProductFileInspection) TableName() string {
	return "product_file_inspection"
}
