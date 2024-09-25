package models

import "gorm.io/datatypes"

type AttributeCategory struct {
    ID               uint           `gorm:"primaryKey;autoIncrement"`
    Code             string         `gorm:"type:varchar(255);not null;uniqueIndex"`
    Name             string         `gorm:"type:varchar(255);not null"`
    RequiredSettings datatypes.JSON `gorm:"type:json;not null"`
}

func (*AttributeCategory) TableName() string {
    return "attribute_category"
}
