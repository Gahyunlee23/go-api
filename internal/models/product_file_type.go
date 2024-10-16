package models

type ProductFileType struct {
	ProductID  uint `gorm:"primaryKey"`
	FileTypeID uint `gorm:"primaryKey"`
	IsDefault  bool `gorm:"default:0;not null" json:"is_default"`
}

func (*ProductFileType) TableName() string {
	return "product_file_type"
}
