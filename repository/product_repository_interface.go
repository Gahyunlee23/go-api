package repository

import (
	"main-admin-api/models"

	"github.com/gin-gonic/gin"
)

type ProductRepositoryInterface interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error)
	GetAll(ctx *gin.Context) ([]models.ProductLite, error)
	Update(product *models.Product) error
	Delete(id uint) error
	ArchiveProduct(id uint) error
}
