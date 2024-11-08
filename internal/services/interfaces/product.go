package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	CreateProduct(product *models.Product, ctx *gin.Context) error
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts(ctx *gin.Context) (*models.ListResponse[models.ProductLite], error)
	UpdateProduct(urlID uint, product *models.Product, ctx *gin.Context) error
	DeleteProduct(id uint) error
	ArchiveProduct(id uint) error
}
