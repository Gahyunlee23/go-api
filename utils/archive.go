package utils

import (
	"encoding/json"
	"log"
	"main-admin-api/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func ArchiveAndDelete(tx *gorm.DB, model interface{ TableName() string }, id uint) error {
	if err := tx.First(model, id).Error; err != nil {
		return err
	}

	jsonData, err := json.Marshal(model)
	if err != nil {
		return err
	}

	log.Println(jsonData)
	println("아니뭔데")

	archiveRecord := models.ArchivedRecord{
		SourceTable:  model.TableName(),
		ArchivedData: datatypes.JSON(jsonData),
	}

	if err := tx.Create(&archiveRecord).Error; err != nil {
		return err
	}

	if err := tx.Unscoped().Delete(model, id).Error; err != nil {
		return err
	}

	return nil
}
