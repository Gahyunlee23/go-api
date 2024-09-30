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
	var productPart models.ProductPart
	if err := r.db.First(&productPart, id).Error; err != nil {
		return nil, err
	}
	return &productPart, nil
}

func (r *ProductPartRepositoryImpl) GetAll(ctx *gin.Context) ([]models.ProductPart, error) {
	var productPart []models.ProductPart
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code", "content_type")).Find(&productPart).Error; err != nil {
		return nil, err
	}
	return productPart, nil
}

func (r *ProductPartRepositoryImpl) Update(productPart *models.ProductPart) error {
	return r.db.Save(productPart).Error
}

func (r *ProductPartRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.ProductPart{}, id).Error
}

func (r *ProductPartRepositoryImpl) Archive(id uint) error {
	var productPart models.ProductPart
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, &productPart, id)
	})
}
