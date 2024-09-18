package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: repository}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	//include business logic, for example, validation
	return s.productRepository.Create(product)
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepository.GetAll()
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.productRepository.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepository.Delete(id)
}
