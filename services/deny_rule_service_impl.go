package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"
	"main-admin-api/utils"

	"github.com/gin-gonic/gin"
)

type DenyRuleServiceImpl struct {
	denyRuleRepository repository.DenyRuleRepositoryInterface
}

func NewDenyRuleServiceImpl(repository repository.DenyRuleRepositoryInterface) DenyRuleServiceInterface {
	return &DenyRuleServiceImpl{denyRuleRepository: repository}
}

func (s *DenyRuleServiceImpl) CreateDenyRule(denyRule *models.DenyRule, ctx *gin.Context) error {
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

func (s *DenyRuleServiceImpl) GetDenyRuleByID(id uint) (*models.DenyRule, error) {
	return s.denyRuleRepository.GetByID(id)
}

func (s *DenyRuleServiceImpl) GetAllDenyRules(ctx *gin.Context) ([]models.DenyRule, error) {
	return s.denyRuleRepository.GetAll(ctx)
}

func (s *DenyRuleServiceImpl) UpdateDenyRule(denyRule *models.DenyRule, ctx *gin.Context) error {
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

func (s *DenyRuleServiceImpl) DeleteDenyRule(id uint) error {
	return s.denyRuleRepository.Delete(id)
}

func (s *DenyRuleServiceImpl) ArchiveDenyRule(id uint) error {
	return s.denyRuleRepository.Archive(id)
}
