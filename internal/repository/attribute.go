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

type attributeRepo struct {
	db *gorm.DB
}

func NewAttributeRepository(db *gorm.DB) repository.AttributeRepository {
	return &attributeRepo{db: db}
}

func (r *attributeRepo) Create(Attribute *models.Attribute) error {
	return r.db.Create(Attribute).Error
}

func (r *attributeRepo) GetByID(id uint) (*models.Attribute, error) {
	attribute := &models.Attribute{ID: id}
	if err := r.db.First(attribute).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Attribute",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch attribute: %w", err)
	}
	return attribute, nil
}

func (r *attributeRepo) GetAll(ctx *gin.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "code", "name")).Find(&attributes).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch attributes: %w", err)
	}
	return attributes, nil
}

func (r *attributeRepo) Update(attribute *models.Attribute) error {
	return r.db.Updates(attribute).Error
}

func (r *attributeRepo) Delete(id uint) error {
	return r.db.Delete(id).Error
}

func (r *attributeRepo) Archive(id uint) error {
	attribute := &models.Attribute{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, attribute, id)
	})
}

func (r *attributeRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.Attribute{}).Scopes(utils.Search(ctx, "id", "code", "name")).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}
