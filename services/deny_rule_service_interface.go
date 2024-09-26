package services

import (
	"main-admin-api/models"

	"github.com/gin-gonic/gin"
)

type DenyRuleServiceInterface interface {
	CreateDenyRule(denyRule *models.DenyRule) error
	GetDenyRuleByID(id uint) (*models.DenyRule, error)
	GetAllDenyRules(ctx *gin.Context) ([]models.DenyRule, error)
	UpdateDenyRule(denyRule *models.DenyRule) error
	DeleteDenyRule(id uint) error
}
