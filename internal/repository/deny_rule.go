package repository

import (
	"errors"
	"fmt"
	"log"
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type denyRuleRepo struct {
	db *gorm.DB
}

func NewDenyRuleRepository(db *gorm.DB) repository.DenyRuleRepository {
	return &denyRuleRepo{db: db}
}

func (r *denyRuleRepo) Create(DenyRule *models.DenyRule) error {
	return r.db.Create(DenyRule).Error
}

func (r *denyRuleRepo) GetByID(id uint) (*models.DenyRule, error) {
	DenyRule := &models.DenyRule{ID: id}

	if err := r.db.First(DenyRule, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "DenyRule",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch deny rule: %w", err)
	}

	return DenyRule, nil
}

func (r *denyRuleRepo) GetAll(ctx *gin.Context) ([]models.DenyRule, error) {
	var denyRules []models.DenyRule

	if err := r.db.Model(&models.DenyRule{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "name", "code")).Find(&denyRules).Error; err != nil {
		return nil, err
	}
	return denyRules, nil
}

func (r *denyRuleRepo) Update(denyRule *models.DenyRule) error {
	log.Printf("Updating DenyRule with ID: %d", denyRule.ID)
	return r.db.Model(denyRule).Updates(denyRule).Error
}

func (r *denyRuleRepo) Delete(id uint) error {
	DenyRule := &models.DenyRule{ID: id}
	return r.db.Model(DenyRule).Delete(id).Error
}

func (r *denyRuleRepo) Archive(id uint) error {
	denyRule := &models.DenyRule{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, denyRule, id)
	})
}

func (r *denyRuleRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.DenyRule{}).Scopes(utils.Search(ctx, "id", "code", "name")).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}