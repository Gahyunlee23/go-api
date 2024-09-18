package services

import "main-admin-api/models"

type ProductServiceInterface interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id uint) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}
