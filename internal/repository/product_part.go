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

type productPartRepo struct {
	db *gorm.DB
}

var productPartColumns = models.SearchSortColumns{
	Search: []string{"id", "name", "code", "content_type"},
	Sort:   []string{"id", "name", "code", "content_type"},
}

func NewProductPartRepository(db *gorm.DB) repository.ProductPartRepository {
	return &productPartRepo{db: db}
}

func (r *productPartRepo) Create(productPart *models.ProductPart) error {
	return r.db.Create(productPart).Error
}

func (r *productPartRepo) GetByID(id uint) (*models.ProductPart, error) {
	ProductPart := &models.ProductPart{ID: id}
	if err := r.db.Debug().Preload("Product").First(ProductPart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "ProductPart",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}
	return ProductPart, nil
}

func (r *productPartRepo) GetAll(ctx *gin.Context) ([]models.ProductPart, error) {
	var productPart []models.ProductPart
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, productPartColumns.Search), utils.Sort(ctx, productPartColumns.Sort)).Find(&productPart).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch product part: %w", err)
	}
	return productPart, nil

}

func (r *productPartRepo) Update(productPart *models.ProductPart) error {
	return r.db.Updates(productPart).Error
}

func (r *productPartRepo) Delete(id uint) error {
	return r.db.Delete(id).Error
}

func (r *productPartRepo) Archive(id uint) error {
	productPart := &models.ProductPart{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, productPart, id)
	})
}

func (r *productPartRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64

	if err := r.db.Model(&models.ProductPart{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, productPartColumns.Search)).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}

	return totalCount, nil
}
