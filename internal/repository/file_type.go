package repository

import (
	"errors"
	"fmt"
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type fileTypeRepo struct {
	db *gorm.DB
}

var fileTypeColumns = models.SearchSortColumns{
	Search: []string{"id", "code", "name"},
	Sort:   []string{"id", "code", "name"},
}

func NewFileTypeRepository(db *gorm.DB) repository.FileTypeRepository {
	return &fileTypeRepo{db: db}
}

func (r *fileTypeRepo) Create(fileType *models.FileType) error {
	return r.db.Create(fileType).Error
}

func (r *fileTypeRepo) GetByID(id uint) (*models.FileType, error) {
	fileType := &models.FileType{ID: id}
	if err := r.db.First(fileType).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "File Type",
				ID:         id,
			}
		}
	}
	return fileType, nil
}

func (r *fileTypeRepo) GetAll(ctx *gin.Context) ([]models.FileType, error) {
	var fileTypes []models.FileType
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, fileInspectionColumns.Search), utils.Sort(ctx, fileTypeColumns.Sort)).Find(&fileTypes).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch file type: %w", err)
	}
	return fileTypes, nil
}

func (r *fileTypeRepo) Update(fileType *models.FileType) error {
	if err := r.db.Updates(fileType).Error; err != nil {
		return fmt.Errorf("failed to update file type: %w", err)
	}
	return nil
}

func (r *fileTypeRepo) Archive(id uint) error {
	fileType := &models.FileType{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, fileType, id)
	})
}
func (r *fileTypeRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.FileType{}).Scopes(utils.Search(ctx, fileTypeColumns.Search)).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to count file type: %w", err)
	}
	return totalCount, nil
}
