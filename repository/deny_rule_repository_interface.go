package repository

import (
	"main-admin-api/models"

	"github.com/gin-gonic/gin"
)

type DenyRuleRepositoryInterface interface {
	Create(denyRule *models.DenyRule) error
	GetByID(id uint) (*models.DenyRule, error)
	GetAll(ctx *gin.Context) ([]models.DenyRule, error)
	Update(denyRule *models.DenyRule) error
	Delete(id uint) error
}
