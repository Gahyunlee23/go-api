package services

import "main-admin-api/models"

type ProductPartServiceInterface interface {
	CreateProductPart(productPart *models.ProductPart) error
	GetProductPartByID(id uint) (*models.ProductPart, error)
	GetAllProductPart() ([]models.ProductPart, error)
	UpdateProductPart(productPart *models.ProductPart) error
	DeleteProductPart(id uint) error
}
