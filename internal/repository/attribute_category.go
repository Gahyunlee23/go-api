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

type attributeCategoryRepo struct {
	db *gorm.DB
}

func NewAttributeCategoryRepository(db *gorm.DB) repository.AttributeCategoryRepository {
	return &attributeCategoryRepo{db: db}
}

func (r *attributeCategoryRepo) Create(AttributeCategory *models.AttributeCategory) error {
	return r.db.Create(AttributeCategory).Error
}

func (r *attributeCategoryRepo) GetAll(ctx *gin.Context) ([]models.AttributeCategory, error) {
	var attributeCategory []models.AttributeCategory
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "code", "name")).Find(&attributeCategory).Error; err != nil {
		return nil, err
	}
	return attributeCategory, nil
}

func (r *attributeCategoryRepo) GetByID(id uint) (*models.AttributeCategory, error) {
	AttributeCategory := &models.AttributeCategory{ID: id}

	if err := r.db.First(AttributeCategory, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Attribute Category",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch attribute category: %w", err)
	}

	return AttributeCategory, nil
}

func (r *attributeCategoryRepo) Update(attributeCategory *models.AttributeCategory) error {
	return r.db.Model(attributeCategory).Updates(attributeCategory).Error
}

func (r *attributeCategoryRepo) Archive(id uint) error {
	attributeCategory := &models.AttributeCategory{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, attributeCategory, id)
	})
}

func (r *attributeCategoryRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.AttributeCategory{}).Scopes(utils.Search(ctx, "id", "code", "name")).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}

func (r *attributeCategoryRepo) GetByCategoryID(ctx *gin.Context, categoryID uint) (*models.AttributeCategory, error) {
	var category models.AttributeCategory

	query := r.db.Preload("Attributes", func(db *gorm.DB) *gorm.DB {
		return db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "code", "name"))
	})

	if err := query.First(&category, categoryID).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
