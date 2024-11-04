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

type proofRepo struct {
	db *gorm.DB
}

var proofColumns = models.SearchSortColumns{
	Search: []string{"id", "code", "name", "description", "price"},
	Sort:   []string{"id", "code", "name", "description", "price", "created_at"},
}

func NewProofRepository(db *gorm.DB) repository.ProofRepository {
	return &proofRepo{db: db}
}

func (r *proofRepo) Create(proof *models.Proof) error {
	return r.db.Create(proof).Error
}

func (r *proofRepo) GetByID(id uint) (*models.Proof, error) {
	proof := &models.Proof{ID: id}
	if err := r.db.First(proof).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{
				EntityType: "Proof",
				ID:         id,
			}
		}
		return nil, fmt.Errorf("failed to fetch proof: %w", err)
	}
	return proof, nil
}

func (r *proofRepo) GetAll(ctx *gin.Context) ([]models.Proof, error) {
	var proofs []models.Proof
	if err := r.db.Scopes(utils.Paginate(ctx), utils.Search(ctx, proofColumns.Search), utils.Sort(ctx, proofColumns.Sort)).Find(&proofs).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch proof: %w", err)
	}
	return proofs, nil
}

func (r *proofRepo) Update(proof *models.Proof) error {
	return r.db.Updates(proof).Error
}

func (r *proofRepo) Archive(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return utils.ArchiveAndDelete(tx, &models.Proof{ID: id}, id)
	})
}

func (r *proofRepo) Count(ctx *gin.Context) (int64, error) {
	var totalCount int64
	if err := r.db.Model(&models.Proof{}).Scopes(utils.Search(ctx, proofColumns.Search)).Count(&totalCount).Error; err != nil {
		return 0, fmt.Errorf("failed to count proof: %w", err)
	}
	return totalCount, nil
}
