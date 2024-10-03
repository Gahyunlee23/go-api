package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type SelectionRuleRepository interface {
	Create(*models.SelectionRule) error
	GetByID(id uint) (*models.SelectionRule, error)
	GetAll(ctx *gin.Context) ([]models.SelectionRule, error)
	Update(*models.SelectionRule) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
