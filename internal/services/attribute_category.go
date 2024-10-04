package services

import (
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"

	"github.com/gin-gonic/gin"
)

type attributeCategoryService struct {
	attributeCategoryRepository repository.AttributeCategoryRepository
	attributeRepository         repository.AttributeRepository
}

func NewAttributeCategoryService(
	attributeCategoryRepo repository.AttributeCategoryRepository,
	attributeRepo repository.AttributeRepository,
) services.AttributeCategoryService {
	return &attributeCategoryService{
		attributeCategoryRepository: attributeCategoryRepo,
		attributeRepository:         attributeRepo,
	}
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

func (s *attributeCategoryService) GetAttributesCategoryByID(id uint, ctx *gin.Context) (*models.DetailResponse[models.AttributeCategory, models.Attribute], error) {
	category, err := s.attributeCategoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	attributes, err := s.attributeRepository.GetByCategoryID(id, ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewDetailResponse(*category, attributes)
	return &response, nil
}
