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

type fileInspectionRepo struct {
	db *gorm.DB
}

func NewFileInspectionRepository(db *gorm.DB) repository.FileInspectionRepository {
	return &fileInspectionRepo{db: db}
}

func (r *fileInspectionRepo) Create(fileInspection *models.FileInspection) error {
	return r.db.Create(fileInspection).Error
}

func (r *fileInspectionRepo) GetByID(id uint) (*models.FileInspection, error) {
	fileInspection := &models.FileInspection{ID: id}
	if err := r.db.First(fileInspection).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "File Inspection",
				ID:         id,
			}
		}
	}
	return fileInspection, nil
}

func (r *fileInspectionRepo) GetAll(ctx *gin.Context) ([]models.FileInspection, error) {
	var fileInspections []models.FileInspection
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name")).Find(&fileInspections).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch file inspection: %w", err)
	}
	return fileInspections, nil
}

func (r *fileInspectionRepo) Update(fileInspection *models.FileInspection) error {
	if err := r.db.Updates(fileInspection).Error; err != nil {
		return fmt.Errorf("failed to update file inspection: %w", err)
	}
	return nil
}

func (r *fileInspectionRepo) Archive(id uint) error {
	fileInspection := &models.FileInspection{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, fileInspection, id)
	})
}
func (r *fileInspectionRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.FileInspection{}).Scopes(utils.Search(ctx, "id", "name")).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to count file inspection: %w", err)
	}
	return totalCount, nil
}
