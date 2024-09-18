package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepositoryInterface
}

func NewProductServiceImpl(repository repository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductServiceImpl{productRepository: repository}
}

func (s *ProductServiceImpl) CreateProduct(product *models.Product) error {
	return s.productRepository.Create(product)
}

func (s *ProductServiceImpl) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductServiceImpl) GetAllProducts() ([]models.Product, error) {
	return s.productRepository.GetAll()
}

func (s *ProductServiceImpl) UpdateProduct(product *models.Product) error {
	return s.productRepository.Update(product)
}

func (s *ProductServiceImpl) DeleteProduct(id uint) error {
	return s.productRepository.Delete(id)
}
