package services

import (
	"main-admin-api/internal/models"
	"main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type denyRuleService struct {
	denyRuleRepository repository.DenyRuleRepository
}

func NewDenyRuleService(repository repository.DenyRuleRepository) services.DenyRuleService {
	return &denyRuleService{denyRuleRepository: repository}
}

func (s *denyRuleService) CreateDenyRule(denyRule *models.DenyRule, ctx *gin.Context) error {
	var err error
	denyRule.Paper, err = utils.MarshalAndAssignJSON(denyRule.Paper, "paper", ctx)
	if err != nil {
		return err
	}

	denyRule.Format, err = utils.MarshalAndAssignJSON(denyRule.Format, "format", ctx)
	if err != nil {
		return err
	}

	denyRule.Pages, err = utils.MarshalAndAssignJSON(denyRule.Pages, "pages", ctx)
	if err != nil {
		return err
	}

	denyRule.Colors, err = utils.MarshalAndAssignJSON(denyRule.Colors, "colors", ctx)
	if err != nil {
		return err
	}

	denyRule.BookBinding, err = utils.MarshalAndAssignJSON(denyRule.BookBinding, "bookBinding", ctx)
	if err != nil {
		return err
	}

	denyRule.Refinement, err = utils.MarshalAndAssignJSON(denyRule.Refinement, "refinement", ctx)
	if err != nil {
		return err
	}

	denyRule.Finishing, err = utils.MarshalAndAssignJSON(denyRule.Finishing, "finishing", ctx)
	if err != nil {
		return err
	}

	if err := s.denyRuleRepository.Create(denyRule); err != nil {
		return err
	}

	return nil
}

func (s *denyRuleService) GetDenyRuleByID(id uint) (*models.DenyRule, error) {
	return s.denyRuleRepository.GetByID(id)
}

func (s *denyRuleService) GetAllDenyRules(ctx *gin.Context) ([]models.DenyRule, error) {
	return s.denyRuleRepository.GetAll(ctx)
}

func (s *denyRuleService) UpdateDenyRule(denyRule *models.DenyRule, ctx *gin.Context) error {
	var err error
	denyRule.Paper, err = utils.MarshalAndAssignJSON(denyRule.Paper, "paper", ctx)
	if err != nil {
		return err
	}

	denyRule.Format, err = utils.MarshalAndAssignJSON(denyRule.Format, "format", ctx)
	if err != nil {
		return err
	}

	denyRule.Pages, err = utils.MarshalAndAssignJSON(denyRule.Pages, "pages", ctx)
	if err != nil {
		return err
	}

	denyRule.Colors, err = utils.MarshalAndAssignJSON(denyRule.Colors, "colors", ctx)
	if err != nil {
		return err
	}

	denyRule.BookBinding, err = utils.MarshalAndAssignJSON(denyRule.BookBinding, "bookBinding", ctx)
	if err != nil {
		return err
	}

	denyRule.Refinement, err = utils.MarshalAndAssignJSON(denyRule.Refinement, "refinement", ctx)
	if err != nil {
		return err
	}

	denyRule.Finishing, err = utils.MarshalAndAssignJSON(denyRule.Finishing, "finishing", ctx)
	if err != nil {
		return err
	}

	if err := s.denyRuleRepository.Update(denyRule); err != nil {
		return err
	}

	return nil
}

func (s *denyRuleService) DeleteDenyRule(id uint) error {
	return s.denyRuleRepository.Delete(id)
}

func (s *denyRuleService) ArchiveDenyRule(id uint) error {
	return s.denyRuleRepository.Archive(id)
}
