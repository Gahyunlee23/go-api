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

type productRepo struct {
	db *gorm.DB
}

var productColumns = models.SearchSortColumns{
	Search: []string{"id", "name", "code", "type"},
	Sort:   []string{"id", "name", "code", "type", "created_at"},
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepo) GetByID(id uint) (*models.Product, error) {
	product := &models.Product{ID: id}

	if err := r.db.First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Product",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch product:	 %w", err)
	}
	return product, nil
}

func (r *productRepo) GetByIDWithPreloads(id uint, preloadFields ...string) (*models.Product, error) {
	product := &models.Product{ID: id}

	query := r.db
	for _, preloadField := range preloadFields {
		query = query.Preload(preloadField)
	}

	if err := query.First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Product",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch product:	 %w", err)
	}
	return product, nil
}

func (r *productRepo) GetAll(ctx *gin.Context) ([]models.ProductLite, error) {
	var products []models.ProductLite

	if err := r.db.Model(&models.ProductLite{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, productColumns.Search), utils.Sort(ctx, productColumns.Sort)).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}
	return products, nil
}

func (r *productRepo) Update(product *models.Product) error {
	return r.db.Model(product).Updates(product).Error
}

func (r *productRepo) Delete(id uint) error {
	return r.db.Delete(id).Error
}

func (r *productRepo) Archive(id uint) error {
	product := &models.Product{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, product, id)
	})
}

func (r *productRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64

	if err := r.db.Model(&models.Product{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, productColumns.Search)).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}
