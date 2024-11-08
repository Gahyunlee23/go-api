package services

import (
	"fmt"
	"main-admin-api/internal/models"
	"main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type attributeService struct {
	attributeRepository repository.AttributeRepository
}

func NewAttributeService(repository repository.AttributeRepository) services.AttributeService {
	return &attributeService{attributeRepository: repository}
}

func (s *attributeService) CreateAttribute(attribute *models.Attribute, ctx *gin.Context) error {
	var err error
	attribute.Settings, err = utils.MarshalAndAssignJSON(attribute.Settings, "settings", ctx)
	if err != nil {
		return err
	}

	if err := s.attributeRepository.Create(attribute); err != nil {
		return err
	}

	return nil
}

func (s *attributeService) GetAttributeByID(id uint) (*models.Attribute, error) {
	return s.attributeRepository.GetByID(id)
}

func (s *attributeService) GetAllAttributes(ctx *gin.Context) (*models.ListResponse[models.Attribute], error) {
	attributes, err := s.attributeRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.attributeRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, attributes)
	return &response, nil
}

func (s *attributeService) UpdateAttribute(urlID uint, attribute *models.Attribute, ctx *gin.Context) error {
	// Verify that URL ID matches the attribute ID
	if err := utils.ValidateID(urlID, attribute.ID); err != nil {
		return fmt.Errorf("ID validation failed: %w", err)
	}

	// Check if the attribute exists
	_, err := utils.ValidateAndFetchEntity(s.attributeRepository, urlID, "Attribute")
	if err != nil {
		return err
	}

	// Marshal and assign JSON for settings
	attribute.Settings, err = utils.MarshalAndAssignJSON(attribute.Settings, "settings", ctx)
	if err != nil {
		return fmt.Errorf("error processing settings: %w", err)
	}

	// Update the attribute
	if err := s.attributeRepository.Update(attribute); err != nil {
		return fmt.Errorf("failed to update attribute: %w", err)
	}

	return nil
}

func (s *attributeService) DeleteAttribute(id uint) error {
	return s.attributeRepository.Delete(id)
}

func (s *attributeService) ArchiveAttribute(id uint) error {
	_, err := utils.ValidateAndFetchEntity[models.Attribute](s.attributeRepository, id, "Fixed Price")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	if err := s.attributeRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to delete fixed price: %w", err)
	}

	return nil
}
