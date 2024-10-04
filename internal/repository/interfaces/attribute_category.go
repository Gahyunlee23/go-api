package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type AttributeCategoryRepository interface {
	Create(AttributeCategory *models.AttributeCategory) error
	GetAll(ctx *gin.Context) ([]models.AttributeCategory, error)
	GetByID(id uint) (*models.AttributeCategory, error)
	Count(ctx *gin.Context) (int64, error)
	Update(AttributeCategory *models.AttributeCategory) error
	Archive(id uint) error
}
