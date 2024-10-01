package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type FixedPriceRepository interface {
	Create(*models.FixedPrice) error
	GetByID(id uint) (*models.FixedPrice, error)
	GetAll(ctx *gin.Context) ([]models.FixedPrice, error)
	Update(*models.FixedPrice) error
	Archive(id uint) error
}
