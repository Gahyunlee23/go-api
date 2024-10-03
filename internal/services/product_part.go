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

type ProductPartServiceImpl struct {
	productPartRepository repository.ProductPartRepository
}

func NewProductPartService(repository repository.ProductPartRepository) services.ProductPartService {
	return &ProductPartServiceImpl{productPartRepository: repository}
}

func (s *ProductPartServiceImpl) CreateProductPart(productPart *models.ProductPart, ctx *gin.Context) error {
	// JSON fields to process
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&productPart.Paper, "paper"},
		{&productPart.Format, "format"},
		{&productPart.Pages, "pages"},
		{&productPart.Colors, "colors"},
		{&productPart.BookBinding, "bookBinding"},
		{&productPart.Refinement, "refinement"},
		{&productPart.Finishing, "finishing"},
		{&productPart.DefaultSelections, "defaultSelections"},
	}

	// Process all JSON fields
	for _, item := range jsonFields {
		if err, _ := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	if err := s.productPartRepository.Create(productPart); err != nil {
		return err
	}

	return nil
}

func (s *ProductPartServiceImpl) GetProductPartByID(id uint) (*models.ProductPart, error) {
	return s.productPartRepository.GetByID(id)
}

func (s *ProductPartServiceImpl) GetAllProductPart(ctx *gin.Context) ([]models.ProductPart, error) {
	return s.productPartRepository.GetAll(ctx)
}

func (s *ProductPartServiceImpl) UpdateProductPart(urlID uint, productPart *models.ProductPart, ctx *gin.Context) error {
	// Verify that URL ID matches the product part ID
	if urlID != productPart.ID {
		return errors.New("product part ID in URL does not match the ID in the request body")
	}

	// Check if the product part exists
	_, err := utils.ValidateAndFetchEntity[models.ProductPart](s.productPartRepository, urlID, "Product Part")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	// JSON fields to process
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&productPart.Paper, "paper"},
		{&productPart.Format, "format"},
		{&productPart.Pages, "pages"},
		{&productPart.Colors, "colors"},
		{&productPart.BookBinding, "bookBinding"},
		{&productPart.Refinement, "refinement"},
		{&productPart.Finishing, "finishing"},
		{&productPart.DefaultSelections, "defaultSelections"},
	}

	// Process all JSON fields
	for _, item := range jsonFields {
		if err, _ := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	// Update the product part
	if err := s.productPartRepository.Update(productPart); err != nil {
		return fmt.Errorf("failed to update product part: %w", err)
	}

	return nil
}

func (s *ProductPartServiceImpl) DeleteProductPart(id uint) error {
	return s.productPartRepository.Delete(id)
}

func (s *ProductPartServiceImpl) ArchiveProductPart(id uint) error {
	_, err := utils.ValidateAndFetchEntity[models.ProductPart](s.productPartRepository, id, "Product Part")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	if err := s.productPartRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to archive fixed price: %w", err)
	}

	return nil
}
