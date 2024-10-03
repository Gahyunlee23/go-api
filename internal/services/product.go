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

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) services.ProductService {
	return &productService{productRepository: repository}
}

func (s *productService) CreateProduct(product *models.Product, ctx *gin.Context) error {
	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&product.RenamingRules, "renaming_rules"},
		{&product.OrderRules, "order_rules"},
		{&product.QuantitiesSelection, "quantities_selection"},
	}

	for _, item := range jsonFields {
		if err, _ := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	if err := s.productRepository.Create(product); err != nil {
		return err
	}

	return nil
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *productService) GetAllProducts(ctx *gin.Context) (*models.ListResponse[models.ProductLite], error) {
	product, err := s.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.productRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, product)
	return &response, nil
}

func (s *productService) UpdateProduct(urlID uint, product *models.Product, ctx *gin.Context) error {
	if urlID != product.ID {
		return errors.New("product ID in URL does not match the ID in the request body")
	}

	_, err := utils.ValidateAndFetchEntity[models.Product](s.productRepository, urlID, "Product")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	jsonFields := []struct {
		field interface{}
		name  string
	}{
		{&product.RenamingRules, "renaming_rules"},
		{&product.OrderRules, "order_rules"},
		{&product.QuantitiesSelection, "quantities_selection"},
	}

	for _, item := range jsonFields {
		if err, _ := utils.MarshalAndAssignJSON(item.field, item.name, ctx); err != nil {
			return fmt.Errorf("error processing %s: %w", item.name, err)
		}
	}

	if err := s.productRepository.Update(product); err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

func (s *productService) DeleteProduct(id uint) error {
	return s.productRepository.Delete(id)
}

func (s *productService) ArchiveProduct(id uint) error {
	_, err := utils.ValidateAndFetchEntity[models.Product](s.productRepository, id, "Product")
	if err != nil {
		return fmt.Errorf("failed to validate fixed price: %w", err)
	}

	if err := s.productRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to archive fixed price: %w", err)
	}

	return nil
}
