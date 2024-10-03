package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type DenyRuleRepository interface {
	Create(denyRule *models.DenyRule) error
	GetByID(id uint) (*models.DenyRule, error)
	GetAll(ctx *gin.Context) ([]models.DenyRule, error)
	Update(denyRule *models.DenyRule) error
	Delete(id uint) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
