package services

import (
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
	var err error

	product.RenamingRules, err = utils.MarshalAndAssignJSON(product.RenamingRules, "renaming_rules", ctx)
	if err != nil {
		return err
	}

	product.OrderRules, err = utils.MarshalAndAssignJSON(product.OrderRules, "order_rules", ctx)
	if err != nil {
		return err
	}

	product.QuantitiesSelection, err = utils.MarshalAndAssignJSON(product.QuantitiesSelection, "quantities_selection", ctx)
	if err != nil {
		return err
	}

	if err := s.productRepository.Create(product); err != nil {
		return err
	}

	return nil
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *productService) GetAllProducts(ctx *gin.Context) ([]models.ProductLite, error) {
	return s.productRepository.GetAll(ctx)
}

func (s *productService) UpdateProduct(product *models.Product, ctx *gin.Context) error {
	var err error

	product.RenamingRules, err = utils.MarshalAndAssignJSON(product.RenamingRules, "renaming_rules", ctx)
	if err != nil {
		return err
	}

	product.OrderRules, err = utils.MarshalAndAssignJSON(product.OrderRules, "order_rules", ctx)
	if err != nil {
		return err
	}

	product.QuantitiesSelection, err = utils.MarshalAndAssignJSON(product.QuantitiesSelection, "quantities_selection", ctx)
	if err != nil {
		return err
	}

	if err := s.productRepository.Update(product); err != nil {
		return err
	}

	return nil
}

func (s *productService) DeleteProduct(id uint) error {
	return s.productRepository.Delete(id)
}

func (s *productService) ArchiveProduct(id uint) error {
	return s.productRepository.Archive(id)
}
