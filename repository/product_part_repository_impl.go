package repository

import (
	"main-admin-api/models"
	"main-admin-api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductPartRepositoryImpl struct {
	db *gorm.DB
}

func NewProductPartRepositoryImpl(db *gorm.DB) *ProductPartRepositoryImpl {
	return &ProductPartRepositoryImpl{db: db}
}

func (r *ProductPartRepositoryImpl) Create(productPart *models.ProductPart) error {
	return r.db.Create(productPart).Error
}

func (r *ProductPartRepositoryImpl) GetByID(id uint) (*models.ProductPart, error) {
	ProductPart := &models.ProductPart{ID: id}
	if err := r.db.Model(ProductPart).First(ProductPart).Error; err != nil {
		return nil, err
	}
	return ProductPart, nil
}

func (r *ProductPartRepositoryImpl) GetAll(ctx *gin.Context) ([]models.ProductPart, error) {
	var productPart []models.ProductPart
	if err := r.db.Model(&models.ProductPart{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code", "content_type")).Find(&productPart).Error; err != nil {
		return nil, err
	}
	return productPart, nil
}

func (r *ProductPartRepositoryImpl) Update(productPart *models.ProductPart) error {
	return r.db.Model(productPart).Updates(productPart).Error
}

func (r *ProductPartRepositoryImpl) Delete(id uint) error {
	productPart := &models.ProductPart{ID: id}
	return r.db.Model(productPart).Delete(id).Error
}

func (r *ProductPartRepositoryImpl) Archive(id uint) error {
	productPart := &models.ProductPart{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, productPart, id)
	})
}
