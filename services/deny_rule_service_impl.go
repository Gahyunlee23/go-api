package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"

	"github.com/gin-gonic/gin"
)

type DenyRuleServiceImpl struct {
	denyRuleRepository repository.DenyRuleRepositoryInterface
}

func NewDenyRuleServiceImpl(repository repository.DenyRuleRepositoryInterface) DenyRuleServiceInterface {
	return &DenyRuleServiceImpl{denyRuleRepository: repository}
}

func (s *DenyRuleServiceImpl) CreateDenyRule(denyRule *models.DenyRule) error {
	return s.denyRuleRepository.Create(denyRule)
}

func (s *DenyRuleServiceImpl) GetDenyRuleByID(id uint) (*models.DenyRule, error) {
	return s.denyRuleRepository.GetByID(id)
}

func (s *DenyRuleServiceImpl) GetAllDenyRules(ctx *gin.Context) ([]models.DenyRule, error) {
	return s.denyRuleRepository.GetAll(ctx)
}

func (s *DenyRuleServiceImpl) UpdateDenyRule(denyRule *models.DenyRule) error {
	return s.denyRuleRepository.Update(denyRule)
}

func (s *DenyRuleServiceImpl) DeleteDenyRule(id uint) error {
	return s.denyRuleRepository.Delete(id)
}
