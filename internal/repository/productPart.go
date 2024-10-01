package repository

import (
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productPartRepo struct {
	db *gorm.DB
}

func NewProductPartRepository(db *gorm.DB) repository.ProductPartRepository {
	return &productPartRepo{db: db}
}

func (r *productPartRepo) Create(productPart *models.ProductPart) error {
	return r.db.Create(productPart).Error
}

func (r *productPartRepo) GetByID(id uint) (*models.ProductPart, error) {
	ProductPart := &models.ProductPart{ID: id}
	if err := r.db.Model(ProductPart).First(ProductPart).Error; err != nil {
		return nil, err
	}
	return ProductPart, nil
}

func (r *productPartRepo) GetAll(ctx *gin.Context) ([]models.ProductPart, error) {
	var productPart []models.ProductPart
	if err := r.db.Model(&models.ProductPart{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code", "content_type")).Find(&productPart).Error; err != nil {
		return nil, err
	}
	return productPart, nil
}

func (r *productPartRepo) Update(productPart *models.ProductPart) error {
	return r.db.Model(productPart).Updates(productPart).Error
}

func (r *productPartRepo) Delete(id uint) error {
	productPart := &models.ProductPart{ID: id}
	return r.db.Model(productPart).Delete(id).Error
}

func (r *productPartRepo) Archive(id uint) error {
	productPart := &models.ProductPart{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, productPart, id)
	})
}
