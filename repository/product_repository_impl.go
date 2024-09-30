package repository

import (
	"main-admin-api/models"
	"main-admin-api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepositoryImpl) GetByID(id uint) (*models.Product, error) {
	Product := &models.Product{ID: id}
	if err := r.db.Model(Product).First(Product).Error; err != nil {
		return nil, err
	}
	return Product, nil
}

func (r *ProductRepositoryImpl) GetAll(ctx *gin.Context) ([]models.ProductLite, error) {
	var products []models.ProductLite
	if err := r.db.Model(&models.Product{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code", "type")).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) Update(product *models.Product) error {
	return r.db.Model(product).Updates(product).Error
}

func (r *ProductRepositoryImpl) Delete(id uint) error {
	product := &models.Product{ID: id}
	return r.db.Model(product).Delete(id).Error
}

func (r *ProductRepositoryImpl) Archive(id uint) error {
	product := &models.Product{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx.Model(product), product, id)
	})
}
