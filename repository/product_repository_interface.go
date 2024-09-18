package repository

import "main-admin-api/models"

type ProductRepositoryInterface interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}
