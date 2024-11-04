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

type denyRuleRepo struct {
	db *gorm.DB
}

var denyRuleColumns = models.SearchSortColumns{
	Search: []string{"id", "name", "code"},
	Sort:   []string{"id", "name", "code", "created_at"},
}

func NewDenyRuleRepository(db *gorm.DB) repository.DenyRuleRepository {
	return &denyRuleRepo{db: db}
}

func (r *denyRuleRepo) Create(DenyRule *models.DenyRule) error {
	return r.db.Create(DenyRule).Error
}

func (r *denyRuleRepo) GetByID(id uint) (*models.DenyRule, error) {
	denyRule := &models.DenyRule{ID: id}

	if err := r.db.First(denyRule).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "DenyRule",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch deny rule: %w", err)
	}

	return denyRule, nil
}

func (r *denyRuleRepo) GetAll(ctx *gin.Context) ([]models.DenyRule, error) {
	var denyRules []models.DenyRule

	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, attributeCategoryColumns.Search), utils.Sort(ctx, attributeCategoryColumns.Sort)).Find(&denyRules).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch deny rules: %w", err)
	}
	return denyRules, nil
}

func (r *denyRuleRepo) Update(denyRule *models.DenyRule) error {
	return r.db.Model(denyRule).Updates(denyRule).Error
}

func (r *denyRuleRepo) Delete(id uint) error {
	return r.db.Delete(id).Error
}

func (r *denyRuleRepo) Archive(id uint) error {
	denyRule := &models.DenyRule{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, denyRule, id)
	})
}

func (r *denyRuleRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.DenyRule{}).Scopes(utils.Search(ctx, attributeCategoryColumns.Search)).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}
