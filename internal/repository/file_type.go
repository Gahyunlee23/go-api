package repository

import (
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"

	"gorm.io/gorm"
)

type fileTypeRepo struct {
	db *gorm.DB
}

func NewFileTypeRepository(db *gorm.DB) repository.FileTypeRepository {
	return &fileTypeRepo{db: db}
}

func (r *fileTypeRepo) Create(fileType *models.FileType) error {
	return r.db.Create(fileType).Error
}

func (r *fileTypeRepo) GetByID(id uint) (*models.FileType, error) {
	fileType := &models.FileType{ID: id}
	if err := r.db.Model(fileType).First(&fileType).Error; err != nil {
		return nil, err
	}
	return fileType, nil
}
