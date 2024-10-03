package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductPartRepository interface {
	Create(productPart *models.ProductPart) error
	GetByID(id uint) (*models.ProductPart, error)
	GetAll(ctx *gin.Context) ([]models.ProductPart, error)
	Update(productPart *models.ProductPart) error
	Delete(id uint) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
