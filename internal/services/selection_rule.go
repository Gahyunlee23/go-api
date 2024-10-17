package services

import (
	"fmt"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type selectionRuleService struct {
	selectionRuleRepository repository.SelectionRuleRepository
}

func NewSelectionRuleService(repository repository.SelectionRuleRepository) services.SelectionRuleService {
	return &selectionRuleService{selectionRuleRepository: repository}
}

func (s *selectionRuleService) CreateSelectionRule(ctx *gin.Context, selectionRule *models.SelectionRule) error {
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&selectionRule.DenyRules, "denyRules"},
		{&selectionRule.Paper, "paper"},
		{&selectionRule.Format, "format"},
		{&selectionRule.Pages, "pages"},
		{&selectionRule.Colors, "colors"},
		{&selectionRule.BookBinding, "bookBinding"},
		{&selectionRule.Refinement, "refinement"},
		{&selectionRule.Finishing, "finishing"},
	}

	for _, item := range jsonFields {
		if _, err := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	if err := s.selectionRuleRepository.Create(selectionRule); err != nil {
		return fmt.Errorf("error creating %s: %w", selectionRule.Name, err)
	}

	return nil
}

func (s *selectionRuleService) GetSelectionRuleByID(id uint) (*models.SelectionRule, error) {
	return s.selectionRuleRepository.GetByID(id)
}

func (s *selectionRuleService) GetAllSelectionRules(ctx *gin.Context) (*models.ListResponse[models.SelectionRule], error) {
	selectionRule, err := s.selectionRuleRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	totalCount, err := s.selectionRuleRepository.Count(ctx)
	if err != nil {
		return nil, err
	}
	response := models.NewListResponse(totalCount, selectionRule)
	return &response, nil
}

func (s *selectionRuleService) UpdateSelectionRule(urlID uint, selectionRule *models.SelectionRule, ctx *gin.Context) error {
	// Verify that URL ID matches the selection rule ID
	if err := utils.ValidateID(urlID, selectionRule.ID); err != nil {
		return fmt.Errorf("ID validation failed: %w", err)
	}

	// Check if the selection rule exists
	if _, err := utils.ValidateAndFetchEntity[models.SelectionRule](s.selectionRuleRepository, urlID, "Selection Rule"); err != nil {
		return fmt.Errorf("failed to validate selection rule: %w", err)
	}

	// JSON fields to process
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&selectionRule.Paper, "paper"},
		{&selectionRule.Format, "format"},
		{&selectionRule.Pages, "pages"},
		{&selectionRule.Colors, "colors"},
		{&selectionRule.BookBinding, "bookBinding"},
		{&selectionRule.Refinement, "refinement"},
		{&selectionRule.Finishing, "finishing"},
	}

	// Process all JSON fields
	for _, item := range jsonFields {
		if _, err := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	// Update the selection rule
	if err := s.selectionRuleRepository.Update(selectionRule); err != nil {
		return fmt.Errorf("failed to update selection rule: %w", err)
	}

	return nil
}

func (s *selectionRuleService) ArchiveSelectionRule(id uint) error {
	if _, err := utils.ValidateAndFetchEntity[models.SelectionRule](s.selectionRuleRepository, id, "Selection Rule"); err != nil {
		return fmt.Errorf("failed to validate selection rule: %w", err)
	}

	if err := s.selectionRuleRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to archive selection rule: %w", err)
	}

	return nil
}
