package repository

import (
	"errors"
	"fmt"
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type fixedPriceRepo struct {
	db *gorm.DB
}

func NewFixedPriceRepository(db *gorm.DB) repository.FixedPriceRepository {
	return &fixedPriceRepo{db: db}
}

func (r *fixedPriceRepo) Create(fixedPrice *models.FixedPrice) error {
	return r.db.Create(fixedPrice).Error
}

func (r *fixedPriceRepo) GetByID(id uint) (*models.FixedPrice, error) {
	FixedPrice := &models.FixedPrice{ID: id}
	if err := r.db.Model(FixedPrice).First(FixedPrice).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Fixed Price",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}
	return FixedPrice, nil
}

func (r *fixedPriceRepo) GetAll(ctx *gin.Context) ([]models.FixedPrice, error) {
	var FixedPrice []models.FixedPrice
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code")).Find(&FixedPrice).Error; err != nil {
		return nil, err
	}
	return FixedPrice, nil
}

func (r *fixedPriceRepo) Update(fixedPrice *models.FixedPrice) error {
	return r.db.Model(fixedPrice).Updates(fixedPrice).Error
}

func (r *fixedPriceRepo) Archive(id uint) error {
	fixedPrice := &models.FixedPrice{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, fixedPrice, id)
	})
}

func (r *fixedPriceRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64

	if err := r.db.Model(&models.FixedPrice{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code")).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}
