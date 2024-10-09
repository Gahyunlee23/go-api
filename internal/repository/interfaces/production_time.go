package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductionTimeRepository interface {
	Create(productionTime *models.ProductionTime) error
	GetAll(ctx *gin.Context) ([]models.ProductionTime, error)
	GetByID(id uint) (*models.ProductionTime, error)
	Update(productionTime *models.ProductionTime) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
