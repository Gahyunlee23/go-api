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

type selectionRuleRepo struct {
	db *gorm.DB
}

var selectionTimeColumns = models.SearchSortColumns{
	Search: []string{"id", "name", "code"},
	Sort:   []string{"id", "name", "code", "created_at"},
}

func NewSelectionRuleRepository(db *gorm.DB) repository.SelectionRuleRepository {
	return &selectionRuleRepo{db: db}
}

func (r *selectionRuleRepo) Create(selectionRule *models.SelectionRule) error {
	return r.db.Create(selectionRule).Error
}

func (r *selectionRuleRepo) GetByID(id uint) (*models.SelectionRule, error) {
	selectionRule := &models.SelectionRule{ID: id}
	if err := r.db.First(selectionRule).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "SelectionRule",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch selection rule: %w", err)
	}
	return selectionRule, nil
}

func (r *selectionRuleRepo) GetAll(ctx *gin.Context) ([]models.SelectionRule, error) {
	var selectionRules []models.SelectionRule
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, selectionTimeColumns.Search), utils.Sort(ctx, selectionTimeColumns.Sort)).Find(&selectionRules).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch selection rules: %w", err)
	}
	return selectionRules, nil
}

func (r *selectionRuleRepo) Update(selectionRule *models.SelectionRule) error {
	return r.db.Debug().Updates(selectionRule).Error
}

func (r *selectionRuleRepo) Archive(id uint) error {
	selectionRule := &models.SelectionRule{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, selectionRule, id)
	})
}

func (r *selectionRuleRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.SelectionRule{}).Scopes(utils.Search(ctx, selectionTimeColumns.Search)).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}
