package repository

import "main-admin-api/models"

type ProductPartRepositoryInterface interface {
	Create(productPart *models.ProductPart) error
	GetByID(id uint) (*models.ProductPart, error)
	GetAll() ([]models.ProductPart, error)
	Update(productPart *models.ProductPart) error
	Delete(id uint) error
}
