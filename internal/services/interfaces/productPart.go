package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductPartService interface {
	CreateProductPart(productPart *models.ProductPart, ctx *gin.Context) error
	GetProductPartByID(id uint) (*models.ProductPart, error)
	GetAllProductPart(ctx *gin.Context) ([]models.ProductPart, error)
	UpdateProductPart(urlID uint, productPart *models.ProductPart, ctx *gin.Context) error
	DeleteProductPart(id uint) error
	ArchiveProductPart(id uint) error
}
