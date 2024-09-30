package services

import (
	"main-admin-api/models"

	"github.com/gin-gonic/gin"
)

type ProductServiceInterface interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts(ctx *gin.Context) ([]models.ProductLite, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	ArchiveProduct(id uint) error
}
