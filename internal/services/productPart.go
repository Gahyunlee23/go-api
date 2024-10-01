package services

import (
	"main-admin-api/internal/models"
	"main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProductPartServiceImpl struct {
	productPartRepository repository.ProductPartRepository
}

func NewProductPartService(repository repository.ProductPartRepository) services.ProductPartService {
	return &ProductPartServiceImpl{productPartRepository: repository}
}

func (s *ProductPartServiceImpl) CreateProductPart(productPart *models.ProductPart, ctx *gin.Context) error {
	var err error
	productPart.Paper, err = utils.MarshalAndAssignJSON(productPart.Paper, "paper", ctx)
	if err != nil {
		return err
	}

	productPart.Format, err = utils.MarshalAndAssignJSON(productPart.Format, "format", ctx)
	if err != nil {
		return err
	}

	productPart.Pages, err = utils.MarshalAndAssignJSON(productPart.Pages, "pages", ctx)
	if err != nil {
		return err
	}

	productPart.Colors, err = utils.MarshalAndAssignJSON(productPart.Colors, "colors", ctx)
	if err != nil {
		return err
	}

	productPart.BookBinding, err = utils.MarshalAndAssignJSON(productPart.BookBinding, "bookBinding", ctx)
	if err != nil {
		return err
	}

	productPart.Refinement, err = utils.MarshalAndAssignJSON(productPart.Refinement, "refinement", ctx)
	if err != nil {
		return err
	}

	productPart.Finishing, err = utils.MarshalAndAssignJSON(productPart.Finishing, "finishing", ctx)
	if err != nil {
		return err
	}

	productPart.DefaultSelections, err = utils.MarshalAndAssignJSON(productPart.DefaultSelections, "defaultSelections", ctx)
	if err != nil {
		return err
	}

	if err := s.productPartRepository.Create(productPart); err != nil {
		return err
	}

	return nil
}

func (s *ProductPartServiceImpl) GetProductPartByID(id uint) (*models.ProductPart, error) {
	return s.productPartRepository.GetByID(id)
}

func (s *ProductPartServiceImpl) GetAllProductPart(ctx *gin.Context) ([]models.ProductPart, error) {
	return s.productPartRepository.GetAll(ctx)
}

func (s *ProductPartServiceImpl) UpdateProductPart(productPart *models.ProductPart, ctx *gin.Context) error {
	var err error
	productPart.Paper, err = utils.MarshalAndAssignJSON(productPart.Paper, "paper", ctx)
	if err != nil {
		return err
	}

	productPart.Format, err = utils.MarshalAndAssignJSON(productPart.Format, "format", ctx)
	if err != nil {
		return err
	}

	productPart.Pages, err = utils.MarshalAndAssignJSON(productPart.Pages, "pages", ctx)
	if err != nil {
		return err
	}

	productPart.Colors, err = utils.MarshalAndAssignJSON(productPart.Colors, "colors", ctx)
	if err != nil {
		return err
	}

	productPart.BookBinding, err = utils.MarshalAndAssignJSON(productPart.BookBinding, "bookBinding", ctx)
	if err != nil {
		return err
	}

	productPart.Refinement, err = utils.MarshalAndAssignJSON(productPart.Refinement, "refinement", ctx)
	if err != nil {
		return err
	}

	productPart.Finishing, err = utils.MarshalAndAssignJSON(productPart.Finishing, "finishing", ctx)
	if err != nil {
		return err
	}

	productPart.DefaultSelections, err = utils.MarshalAndAssignJSON(productPart.DefaultSelections, "defaultSelections", ctx)
	if err != nil {
		return err
	}

	if err := s.productPartRepository.Update(productPart); err != nil {
		return err
	}

	return nil
}

func (s *ProductPartServiceImpl) DeleteProductPart(id uint) error {
	return s.productPartRepository.Delete(id)
}

func (s *ProductPartServiceImpl) ArchiveProductPart(id uint) error {
	return s.productPartRepository.Archive(id)
}
