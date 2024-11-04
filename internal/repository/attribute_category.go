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

var attributeCategoryColumns = models.SearchSortColumns{
	Search: []string{"id", "code", "name"},
	Sort:   []string{"id", "code", "name", "created_at"},
}

func NewAttributeCategoryRepository(db *gorm.DB) repository.AttributeCategoryRepository {
	return &attributeCategoryRepo{db: db}
}

func (r *attributeCategoryRepo) Create(AttributeCategory *models.AttributeCategory) error {
	return r.db.Create(AttributeCategory).Error
}

func (r *attributeCategoryRepo) GetAll(ctx *gin.Context) ([]models.AttributeCategory, error) {
	var attributeCategory []models.AttributeCategory
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, attributeColumns.Search), utils.Sort(ctx, attributeCategoryColumns.Sort)).Find(&attributeCategory).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch attribute categories: %w", err)
	}
	return attributeCategory, nil
}

func (r *attributeCategoryRepo) GetByID(id uint) (*models.AttributeCategory, error) {
	attributeCategory := &models.AttributeCategory{ID: id}

	if err := r.db.First(attributeCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Attribute Category",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch attribute category: %w", err)
	}

	return attributeCategory, nil
}

func (r *attributeCategoryRepo) Update(attributeCategory *models.AttributeCategory) error {
	return r.db.Updates(attributeCategory).Error
}

func (r *attributeCategoryRepo) Archive(id uint) error {
	attributeCategory := &models.AttributeCategory{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, attributeCategory, id)
	})
}

func (r *attributeCategoryRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.AttributeCategory{}).Scopes(utils.Search(ctx, attributeCategoryColumns.Search)).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}

func (r *attributeCategoryRepo) GetByCategoryID(ctx *gin.Context, categoryID uint) (*models.AttributeCategory, error) {
	var category models.AttributeCategory

	query := r.db.Preload("Attributes", func(db *gorm.DB) *gorm.DB {
		return db.Scopes(utils.Paginate(ctx), utils.Search(ctx, attributeCategoryColumns.Search))
	})

	if err := query.First(&category, categoryID).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
