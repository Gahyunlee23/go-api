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

type fixedPriceService struct {
	fixedPriceRepository repository.FixedPriceRepository
}

func NewFixedPriceService(repository repository.FixedPriceRepository) services.FixedPriceService {
	return &fixedPriceService{fixedPriceRepository: repository}
}

func (s *fixedPriceService) CreateFixedPrice(ctx *gin.Context, fixedPrice *models.FixedPrice) error {
	// JSON fields to process
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&fixedPrice.Paper, "paper"},
		{&fixedPrice.Format, "format"},
		{&fixedPrice.Pages, "pages"},
		{&fixedPrice.Colors, "colors"},
		{&fixedPrice.BookBinding, "bookBinding"},
		{&fixedPrice.Refinement, "refinement"},
		{&fixedPrice.Finishing, "finishing"},
	}

	// Process all JSON fields
	for _, item := range jsonFields {
		if _, err := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	if err := s.fixedPriceRepository.Create(fixedPrice); err != nil {
		return err
	}

	return nil
}

func (s *fixedPriceService) GetFixedPriceByID(id uint) (*models.FixedPrice, error) {
	return s.fixedPriceRepository.GetByID(id)
}

func (s *fixedPriceService) GetAllFixedPrices(ctx *gin.Context) (*models.ListResponse[models.FixedPrice], error) {
	fixedPrice, err := s.fixedPriceRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.fixedPriceRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, fixedPrice)
	return &response, nil
}

func (s *fixedPriceService) UpdateFixedPrice(urlID uint, fixedPrice *models.FixedPrice, ctx *gin.Context) error {
	// Verify that URL ID matches the fixed price ID
	if urlID != fixedPrice.ID {
		return errors.New("fixed price ID in URL does not match the ID in the request body")
	}

	// Check if the fixed price exists
	_, err := utils.ValidateAndFetchEntity[models.FixedPrice](s.fixedPriceRepository, urlID, "Fixed Price")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	// JSON fields to process
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&fixedPrice.Paper, "paper"},
		{&fixedPrice.Format, "format"},
		{&fixedPrice.Pages, "pages"},
		{&fixedPrice.Colors, "colors"},
		{&fixedPrice.BookBinding, "bookBinding"},
		{&fixedPrice.Refinement, "refinement"},
		{&fixedPrice.Finishing, "finishing"},
	}

	// Process all JSON fields
	for _, item := range jsonFields {
		if _, err := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	// Update the fixed price
	if err := s.fixedPriceRepository.Update(fixedPrice); err != nil {
		return fmt.Errorf("failed to update fixed price: %w", err)
	}

	return nil
}

func (s *fixedPriceService) ArchiveFixedPrice(id uint) error {
	_, err := utils.ValidateAndFetchEntity[models.FixedPrice](s.fixedPriceRepository, id, "Fixed Price")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	if err := s.fixedPriceRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to archive fixed price: %w", err)
	}

	return nil
}
