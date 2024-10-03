package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type SelectionRuleService interface {
	CreateSelectionRule(ctx *gin.Context, rule *models.SelectionRule) error
	GetSelectionRuleByID(id uint) (*models.SelectionRule, error)
	GetAllSelectionRules(ctx *gin.Context) (*models.ListResponse[models.SelectionRule], error)
	UpdateSelectionRule(id uint, rule *models.SelectionRule, ctx *gin.Context) error
	ArchiveSelectionRule(id uint) error
}
