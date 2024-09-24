package repository

import "main-admin-api/models"

type DenyRuleRepositoryInterface interface {
	Create(denyRule *models.DenyRule) error
	GetByID(id uint) (*models.DenyRule, error)
	GetAll() ([]models.DenyRule, error)
	Update(denyRule *models.DenyRule) error
	Delete(id uint) error
}
