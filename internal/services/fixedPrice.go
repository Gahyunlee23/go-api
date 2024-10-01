package services

import (
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type fixedPriceService struct {
	fixedPriceRepository repository.FixedPriceRepository
}

func NewFixedPriceService(repository repository.FixedPriceRepository) services.FixedPriceService {
	return &fixedPriceService{fixedPriceRepository: repository}
}

func (s *fixedPriceService) CreateFixedPrice(ctx *gin.Context, fixedPrice *models.FixedPrice) error {
	var err error
	fixedPrice.Paper, err = utils.MarshalAndAssignJSON(fixedPrice.Paper, "paper", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Format, err = utils.MarshalAndAssignJSON(fixedPrice.Format, "format", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Pages, err = utils.MarshalAndAssignJSON(fixedPrice.Pages, "pages", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Colors, err = utils.MarshalAndAssignJSON(fixedPrice.Colors, "colors", ctx)
	if err != nil {
		return err
	}

	fixedPrice.BookBinding, err = utils.MarshalAndAssignJSON(fixedPrice.BookBinding, "bookBinding", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Refinement, err = utils.MarshalAndAssignJSON(fixedPrice.Refinement, "refinement", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Finishing, err = utils.MarshalAndAssignJSON(fixedPrice.Finishing, "finishing", ctx)
	if err != nil {
		return err
	}

	if err := s.fixedPriceRepository.Create(fixedPrice); err != nil {
		return err
	}

	return nil
}

func (s *fixedPriceService) GetFixedPriceByID(id uint) (*models.FixedPrice, error) {
	return s.fixedPriceRepository.GetByID(id)
}

func (s *fixedPriceService) GetAllFixedPrices(ctx *gin.Context) ([]models.FixedPrice, error) {
	return s.fixedPriceRepository.GetAll(ctx)
}

func (s *fixedPriceService) UpdateFixedPrice(ctx *gin.Context, fixedPrice *models.FixedPrice) error {
	var err error
	fixedPrice.Paper, err = utils.MarshalAndAssignJSON(fixedPrice.Paper, "paper", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Format, err = utils.MarshalAndAssignJSON(fixedPrice.Format, "format", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Pages, err = utils.MarshalAndAssignJSON(fixedPrice.Pages, "pages", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Colors, err = utils.MarshalAndAssignJSON(fixedPrice.Colors, "colors", ctx)
	if err != nil {
		return err
	}

	fixedPrice.BookBinding, err = utils.MarshalAndAssignJSON(fixedPrice.BookBinding, "bookBinding", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Refinement, err = utils.MarshalAndAssignJSON(fixedPrice.Refinement, "refinement", ctx)
	if err != nil {
		return err
	}

	fixedPrice.Finishing, err = utils.MarshalAndAssignJSON(fixedPrice.Finishing, "finishing", ctx)
	if err != nil {
		return err
	}

	if err := s.fixedPriceRepository.Update(fixedPrice); err != nil {
		return err
	}

	return nil
}

func (s *fixedPriceService) ArchiveFixedPrice(id uint) error {
	return s.fixedPriceRepository.Archive(id)
}
