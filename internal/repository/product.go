package repository

import (
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepo) GetByID(id uint) (*models.Product, error) {
	Product := &models.Product{ID: id}
	if err := r.db.Model(Product).First(Product).Error; err != nil {
		return nil, err
	}
	return Product, nil
}

func (r *productRepo) GetAll(ctx *gin.Context) ([]models.ProductLite, error) {
	var products []models.ProductLite
	if err := r.db.Model(&models.Product{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code", "type")).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Update(product *models.Product) error {
	return r.db.Model(product).Updates(product).Error
}

func (r *productRepo) Delete(id uint) error {
	product := &models.Product{ID: id}
	return r.db.Model(product).Delete(id).Error
}

func (r *productRepo) Archive(id uint) error {
	product := &models.Product{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, product, id)
	})
}
