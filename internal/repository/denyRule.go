package repository

import (
	"log"
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
	if err := r.db.Model(DenyRule).First(DenyRule).Error; err != nil {
		return nil, err
	}
	return DenyRule, nil
}

func (r *denyRuleRepo) GetAll(ctx *gin.Context) ([]models.DenyRule, error) {
	var denyRules []models.DenyRule

	if err := r.db.Model(&models.DenyRule{}).Scopes(utils.Paginate(ctx), utils.Search(ctx, "name", "code")).Find(&denyRules).Error; err != nil {
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
