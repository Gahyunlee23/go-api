package services

import (
	"errors"
	"fmt"
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

func (s *denyRuleService) GetAllDenyRules(ctx *gin.Context) (*models.ListResponse[models.DenyRule], error) {
	denyRules, err := s.denyRuleRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.denyRuleRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, denyRules)
	return &response, nil
}

func (s *denyRuleService) UpdateDenyRule(urlID uint, denyRule *models.DenyRule, ctx *gin.Context) error {
	// Verify that URL ID matches the deny rule ID
	if urlID != denyRule.ID {
		return errors.New("deny rule ID in URL does not match the ID in the request body")
	}

	// Check if the deny rule exists
	_, err := utils.ValidateAndFetchEntity(s.denyRuleRepository, urlID, "Deny Rule")
	if err != nil {
		return err
	}

	// JSON fields to process
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&denyRule.Paper, "paper"},
		{&denyRule.Format, "format"},
		{&denyRule.Pages, "pages"},
		{&denyRule.Colors, "colors"},
		{&denyRule.BookBinding, "bookBinding"},
		{&denyRule.Refinement, "refinement"},
		{&denyRule.Finishing, "finishing"},
	}

	// Process all JSON fields
	for _, item := range jsonFields {
		if err, _ := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	// Update the deny rule
	if err := s.denyRuleRepository.Update(denyRule); err != nil {
		return fmt.Errorf("failed to update deny rule: %w", err)
	}

	return nil
}

func (s *denyRuleService) DeleteDenyRule(id uint) error {
	return s.denyRuleRepository.Delete(id)
}

func (s *denyRuleService) ArchiveDenyRule(id uint) error {
	_, err := utils.ValidateAndFetchEntity[models.DenyRule](s.denyRuleRepository, id, "Deny Rule")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	if err := s.denyRuleRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to archive fixed price: %w", err)
	}

	return nil

}
