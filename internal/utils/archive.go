package utils

import (
	"encoding/json"
	"main-admin-api/internal/models"

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
