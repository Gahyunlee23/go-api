package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"
	"main-admin-api/utils"

	"github.com/gin-gonic/gin"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepositoryInterface
}

func NewProductServiceImpl(repository repository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductServiceImpl{productRepository: repository}
}

func (s *ProductServiceImpl) CreateProduct(product *models.Product, ctx *gin.Context) error {
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

func (s *ProductServiceImpl) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductServiceImpl) GetAllProducts(ctx *gin.Context) ([]models.ProductLite, error) {
	return s.productRepository.GetAll(ctx)
}

func (s *ProductServiceImpl) UpdateProduct(product *models.Product, ctx *gin.Context) error {
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

func (s *ProductServiceImpl) DeleteProduct(id uint) error {
	return s.productRepository.Delete(id)
}

func (s *ProductServiceImpl) ArchiveProduct(id uint) error {
	return s.productRepository.Archive(id)
}
