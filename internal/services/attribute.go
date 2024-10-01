package services

import (
	"errors"
	"fmt"
	"main-admin-api/internal/models"
	"main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (s *attributeService) GetAllAttributes(ctx *gin.Context) ([]models.Attribute, error) {
	return s.attributeRepository.GetAll(ctx)
}

func (s *attributeService) UpdateAttribute(urlID uint, attribute *models.Attribute, ctx *gin.Context) error {
	// Verify that URL ID matches the attribute ID
	if urlID != attribute.ID {
		return errors.New("attribute ID in URL does not match the ID in the request body")
	}

	// Check if the attribute exists
	_, err := s.attributeRepository.GetByID(urlID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("attribute not found")
		}
		return fmt.Errorf("error checking existing attribute: %w", err)
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
	return s.attributeRepository.Archive(id)
}
