package services

import (
	"main-admin-api/internal/repository"
	services "main-admin-api/internal/services/interfaces"

	"gorm.io/gorm"
)

type ServiceFactory struct {
	db *gorm.DB
}

func NewServiceFactory(db *gorm.DB) *ServiceFactory {
	return &ServiceFactory{db: db}
}

func (f *ServiceFactory) CreateProductService() services.ProductService {
	productRepo := repository.NewProductRepository(f.db)
	return NewProductService(productRepo)
}

func (f *ServiceFactory) CreateProductPartService() services.ProductPartService {
	productPartRepo := repository.NewProductPartRepository(f.db)
	return NewProductPartService(productPartRepo)
}

func (f *ServiceFactory) CreateDenyRuleService() services.DenyRuleService {
	denyRuleRepo := repository.NewDenyRuleRepository(f.db)
	return NewDenyRuleService(denyRuleRepo)
}

func (f *ServiceFactory) CreateAttributeService() services.AttributeService {
	attributeRepo := repository.NewAttributeRepository(f.db)
	return NewAttributeService(attributeRepo)
}

func (f *ServiceFactory) CreateFixedPriceService() services.FixedPriceService {
	fixedPriceRepo := repository.NewFixedPriceRepository(f.db)
	return NewFixedPriceService(fixedPriceRepo)
}

func (f *ServiceFactory) CreateSelectionRuleService() services.SelectionRuleService {
	selectionRuleRepo := repository.NewSelectionRuleRepository(f.db)
	return NewSelectionRuleService(selectionRuleRepo)
}

func (f *ServiceFactory) CreateAttributeCategoryService() services.AttributeCategoryService {
	attributeCategoryRepo := repository.NewAttributeCategoryRepository(f.db)
	return NewAttributeCategoryService(attributeCategoryRepo)
}

func (f *ServiceFactory) CreateProductionTimeService() services.ProductionTimeService {
	productionTimeRepo := repository.NewProductionTimeRepository(f.db)
	return NewProductionTimeService(productionTimeRepo)
}

func (f *ServiceFactory) CreateProofService() services.ProofService {
	proofRepo := repository.NewProofRepository(f.db)
	return NewProofService(proofRepo)
}

func (f *ServiceFactory) CreateFileTypeService() services.FileTypeService {
	fileTypeRepo := repository.NewFileTypeRepository(f.db)
	return NewFileTypeService(fileTypeRepo)
}

func (f *ServiceFactory) CreateFileInspectionService() services.FileInspectionService {
	fileInspectionRepo := repository.NewFileInspectionRepository(f.db)
	return NewFileInspectionService(fileInspectionRepo)
}
