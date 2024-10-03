package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error)
	GetAll(ctx *gin.Context) ([]models.ProductLite, error)
	Update(product *models.Product) error
	Delete(id uint) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
