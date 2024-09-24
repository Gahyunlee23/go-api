package services

import "main-admin-api/models"

type DenyRuleServiceInterface interface {
	CreateDenyRule(denyRule *models.DenyRule) error
	GetDenyRuleByID(id uint) (*models.DenyRule, error)
	GetAllDenyRules() ([]models.DenyRule, error)
	UpdateDenyRule(denyRule *models.DenyRule) error
	DeleteDenyRule(id uint) error
}
