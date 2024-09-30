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
	DenyRule := &models.DenyRule{ID: id}
	if err := r.db.Model(DenyRule).First(DenyRule).Error; err != nil {
		return nil, err
	}
	return DenyRule, nil
}

func (r *DenyRuleRepositoryImpl) GetAll(ctx *gin.Context) ([]models.DenyRule, error) {
	var denyRules []models.DenyRule

	if err := r.db.Model(&models.DenyRule{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "name", "code")).Find(&denyRules).Error; err != nil {
		return nil, err
	}
	return denyRules, nil
}

func (r *DenyRuleRepositoryImpl) Update(DenyRule *models.DenyRule) error {
	return r.db.Save(DenyRule).Error
}

func (r *DenyRuleRepositoryImpl) Delete(id uint) error {
	DenyRule := &models.DenyRule{ID: id}
	return r.db.Model(DenyRule).Delete(id).Error
}

func (r *DenyRuleRepositoryImpl) Archive(id uint) error {
	denyRule := &models.DenyRule{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx.Model(denyRule), denyRule, id)
	})
}
