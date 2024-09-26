package services

import (
	"main-admin-api/repository"

	"gorm.io/gorm"
)

type ServiceFactory struct {
	db *gorm.DB
}

func NewServiceFactory(db *gorm.DB) *ServiceFactory {
	return &ServiceFactory{db: db}
}

func (f *ServiceFactory) CreateProductService() ProductServiceInterface {
	productRepo := repository.NewProductRepositoryImpl(f.db)
	return NewProductServiceImpl(productRepo)
}

func (f *ServiceFactory) CreateProductPartService() ProductPartServiceInterface {
	productPartRepo := repository.NewProductPartRepositoryImpl(f.db)
	return NewProductPartServiceImpl(productPartRepo)
}

func (f *ServiceFactory) CreateDenyRuleService() DenyRuleServiceInterface {
	denyRuleRepo := repository.NewDenyRuleRepositoryImpl(f.db)
	return NewDenyRuleServiceImpl(denyRuleRepo)
}

func (f *ServiceFactory) CreateAttributeService() AttributeServiceInterface {
	attributeRepo := repository.NewAttributeRepositoryImpl(f.db)
	return NewAttributeServiceImpl(attributeRepo)
}
