package services

import (
	"main-admin-api/models"
	"main-admin-api/repository"
)

type ProductPartServiceImpl struct {
	productPartRepository repository.ProductPartRepositoryInterface
}

func NewProductPartServiceImpl(repository repository.ProductPartRepositoryInterface) ProductPartServiceInterface {
	return &ProductPartServiceImpl{productPartRepository: repository}
}

func (s *ProductPartServiceImpl) CreateProductPart(productPart *models.ProductPart) error {
	return s.productPartRepository.Create(productPart)
}

func (s *ProductPartServiceImpl) GetProductPartByID(id uint) (*models.ProductPart, error) {
	return s.productPartRepository.GetByID(id)
}

func (s *ProductPartServiceImpl) GetAllProductPart() ([]models.ProductPart, error) {
	return s.productPartRepository.GetAll()
}

func (s *ProductPartServiceImpl) UpdateProductPart(productPart *models.ProductPart) error {
	return s.productPartRepository.Update(productPart)
}

func (s *ProductPartServiceImpl) DeleteProductPart(id uint) error {
	return s.productPartRepository.Delete(id)
}
