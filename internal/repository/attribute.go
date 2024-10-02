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
	Attribute := &models.Attribute{ID: id}
	if err := r.db.First(Attribute, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Attribute",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch attribute: %w", err)
	}
	return Attribute, nil
}

func (r *attributeRepo) GetAll(ctx *gin.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	if err := r.db.Model(&models.Attribute{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "code", "name")).Find(&attributes).Error; err != nil {
		return nil, err
	}
	return attributes, nil
}

func (r *attributeRepo) Update(Attribute *models.Attribute) error {
	return r.db.Model(Attribute).Updates(Attribute).Error
}

func (r *attributeRepo) Delete(id uint) error {
	attribute := &models.Attribute{ID: id}
	return r.db.Model(attribute).Delete(id).Error
}

func (r *attributeRepo) Archive(id uint) error {
	attribute := &models.Attribute{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, attribute, id)
	})
}