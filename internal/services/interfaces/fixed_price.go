package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type FixedPriceService interface {
	CreateFixedPrice(ctx *gin.Context, fixedPrice *models.FixedPrice) error
	GetFixedPriceByID(id uint) (*models.FixedPrice, error)
	GetAllFixedPrices(ctx *gin.Context) (*models.ListResponse[models.FixedPrice], error)
	UpdateFixedPrice(urlID uint, fixedPrice *models.FixedPrice, ctx *gin.Context) error
	ArchiveFixedPrice(id uint) error
}
