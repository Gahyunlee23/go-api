package services

import (
	"fmt"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type attributeCategoryService struct {
	attributeCategoryRepository repository.AttributeCategoryRepository
}

func NewAttributeCategoryService(attributeCategoryRepo repository.AttributeCategoryRepository) services.AttributeCategoryService {
	return &attributeCategoryService{attributeCategoryRepository: attributeCategoryRepo}
}

func (s *attributeCategoryService) GetAllAttributesCategories(ctx *gin.Context) (*models.ListResponse[models.AttributeCategory], error) {
	attributeCategory, err := s.attributeCategoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	totalCount, err := s.attributeCategoryRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, attributeCategory)
	return &response, nil
}

func (s *attributeCategoryService) GetAttributesCategoryByID(id uint, ctx *gin.Context) (*models.AttributeCategory, error) {
	attributeCategory, err := s.attributeCategoryRepository.GetByCategoryID(ctx, id)
	if err != nil {
		return nil, err
	}
	return attributeCategory, nil
}

func (s *attributeCategoryService) CreateAttributeCategory(attributeCategory *models.AttributeCategory) error {
	return s.attributeCategoryRepository.Create(attributeCategory)
}

func (s *attributeCategoryService) UpdateAttributeCategory(urlID uint, attributeCategory *models.AttributeCategory) error {
	// Verify that URL ID matches the attribute ID
	if err := utils.ValidateID(urlID, attributeCategory.ID); err != nil {
		return fmt.Errorf("ID validation failed: %w", err)
	}

	_, err := utils.ValidateAndFetchEntity[models.AttributeCategory](s.attributeCategoryRepository, urlID, "Attribute Category")
	if err != nil {
		return fmt.Errorf("failed to validate attribute category: %w", err)
	}

	if err := s.attributeCategoryRepository.Update(attributeCategory); err != nil {
		return fmt.Errorf("failed to update attribute category: %w", err)
	}
	return nil
}

func (s *attributeCategoryService) ArchiveAttributeCategory(id uint) error {
	_, err := utils.ValidateAndFetchEntity[models.AttributeCategory](s.attributeCategoryRepository, id, "Attribute Category")
	if err != nil {
		return fmt.Errorf("failed to validate attribute category: %w", err)
	}

	if err := s.attributeCategoryRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to delete attribute category: %w", err)
	}

	return nil
}
