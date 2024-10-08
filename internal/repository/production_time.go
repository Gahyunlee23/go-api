package repository

import (
	"fmt"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productionTimeRepo struct {
	db *gorm.DB
}

func NewProductionTimeRepository(db *gorm.DB) repository.ProductionTimeRepository {
	return &productionTimeRepo{db: db}
}

func (r *productionTimeRepo) Create(productionTime *models.ProductionTime) error {
	if err := r.db.Create(productionTime).Error; err != nil {
		return err
	}
	return nil
}

func (r *productionTimeRepo) GetAll(ctx *gin.Context) ([]models.ProductionTime, error) {
	var productionTime []models.ProductionTime
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, "id", "code", "name", "time")).Find(&productionTime).Error; err != nil {
		return nil, err
	}
	return productionTime, nil
}

func (r *productionTimeRepo) GetByID(id uint) (*models.ProductionTime, error) {
	productionTime := &models.ProductionTime{ID: id}
	if err := r.db.First(&productionTime, id).Error; err != nil {
		return nil, err
	}
	return productionTime, nil
}

func (r *productionTimeRepo) Update(id uint, productionTime *models.ProductionTime) error {
	if err := r.db.Model(&models.ProductionTime{ID: id}).Updates(productionTime).Error; err != nil {
		return err
	}
	return nil
}

func (r *productionTimeRepo) Archive(id uint) error {
	productionTime := &models.ProductionTime{ID: id}
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, productionTime, id)
	})
}

func (r *productionTimeRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.ProductionTime{}).Scopes(utils.Search(ctx, "id", "name", "code", "time")).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to fetch count: %w", err)
	}
	return totalCount, nil
}
