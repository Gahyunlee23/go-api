package repository

import (
	"main-admin-api/models"
	"main-admin-api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DenyRuleRepositoryImpl struct {
	db *gorm.DB
}

func NewDenyRuleRepositoryImpl(db *gorm.DB) *DenyRuleRepositoryImpl {
	return &DenyRuleRepositoryImpl{db: db}
}

func (r *DenyRuleRepositoryImpl) Create(DenyRule *models.DenyRule) error {
	return r.db.Create(DenyRule).Error
}

func (r *DenyRuleRepositoryImpl) GetByID(id uint) (*models.DenyRule, error) {
	var denyRule models.DenyRule
	if err := r.db.First(&denyRule, id).Error; err != nil {
		return nil, err
	}
	return &denyRule, nil
}

func (r *DenyRuleRepositoryImpl) GetAll(ctx *gin.Context) ([]models.DenyRule, error) {
	var denyRules []models.DenyRule

	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "name", "code")).Find(&denyRules).Error; err != nil {
		return nil, err
	}
	return denyRules, nil
}

func (r *DenyRuleRepositoryImpl) Update(DenyRule *models.DenyRule) error {
	return r.db.Save(DenyRule).Error
}

func (r *DenyRuleRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.DenyRule{}, id).Error
}
