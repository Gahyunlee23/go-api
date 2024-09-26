package services

import (
	"main-admin-api/models"

	"github.com/gin-gonic/gin"
)

type ProductPartServiceInterface interface {
	CreateProductPart(productPart *models.ProductPart) error
	GetProductPartByID(id uint) (*models.ProductPart, error)
	GetAllProductPart(ctx *gin.Context) ([]models.ProductPart, error)
	UpdateProductPart(productPart *models.ProductPart) error
	DeleteProductPart(id uint) error
}
