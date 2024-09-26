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
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetAll(ctx *gin.Context) ([]models.ProductLite, error) {
	var products []models.ProductLite
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code", "type")).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
