package models

type ProductProof struct {
	ProductID uint `gorm:"primaryKey"`
	ProofID   uint `gorm:"primaryKey"`
	IsDefault bool `gorm:"default:0;not null" json:"is_default"`
}

func (*ProductProof) TableName() string {
	return "product_proof"
}
