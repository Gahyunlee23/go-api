package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type DenyRuleService interface {
	CreateDenyRule(denyRule *models.DenyRule, ctx *gin.Context) error
	GetDenyRuleByID(id uint) (*models.DenyRule, error)
	GetAllDenyRules(ctx *gin.Context) (*models.ListResponse[models.DenyRule], error)
	UpdateDenyRule(id uint, denyRule *models.DenyRule, ctx *gin.Context) error
	DeleteDenyRule(id uint) error
	ArchiveDenyRule(id uint) error
}
