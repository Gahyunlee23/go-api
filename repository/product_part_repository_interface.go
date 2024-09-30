package repository

import (
	"main-admin-api/models"

	"github.com/gin-gonic/gin"
)

type ProductPartRepositoryInterface interface {
	Create(productPart *models.ProductPart) error
	GetByID(id uint) (*models.ProductPart, error)
	GetAll(ctx *gin.Context) ([]models.ProductPart, error)
	Update(productPart *models.ProductPart) error
	Delete(id uint) error
	Archive(id uint) error
}
