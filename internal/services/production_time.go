package services

import (
	"errors"
	"fmt"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type productionTimeService struct {
	productionTimeRepository repository.ProductionTimeRepository
}

func NewProductionTimeService(repo repository.ProductionTimeRepository) services.ProductionTimeService {
	return &productionTimeService{productionTimeRepository: repo}
}

func (s *productionTimeService) CreateProductionTime(productionTime *models.ProductionTime) error {
	return s.productionTimeRepository.Create(productionTime)
}

func (s *productionTimeService) GetAllProductionTimes(ctx *gin.Context) (*models.ListResponse[models.ProductionTime], error) {
	productionTimes, err := s.productionTimeRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	totalCount, err := s.productionTimeRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, productionTimes)
	return &response, nil
}

func (s *productionTimeService) GetProductionTimeByID(id uint) (*models.ProductionTime, error) {
	return s.productionTimeRepository.GetByID(id)
}

func (s *productionTimeService) UpdateProductionTime(urlID uint, productionTime *models.ProductionTime) error {
	if urlID != productionTime.ID {
		return errors.New("production time ID in URL does not match the ID in the request body")
	}

	if _, err := utils.ValidateAndFetchEntity(s.productionTimeRepository, urlID, "production time"); err != nil {
		return fmt.Errorf("failed to validate production time: %w", err)
	}

	if err := s.productionTimeRepository.Update(productionTime); err != nil {
		return fmt.Errorf("failed to update production time: %w", err)
	}

	return nil
}

func (s *productionTimeService) ArchiveProductionTime(id uint) error {
	return s.productionTimeRepository.Archive(id)
}
