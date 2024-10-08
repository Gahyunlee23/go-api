package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProductionTimeService interface {
	CreateProductionTime(productionTime *models.ProductionTime) error
	GetAllProductionTimes(ctx *gin.Context) (*models.ListResponse[models.ProductionTime], error)
	GetProductionTimeByID(id uint) (*models.ProductionTime, error)
	UpdateProductionTime(urlID uint, productionTime *models.ProductionTime) error
	ArchiveProductionTime(id uint) error
}
