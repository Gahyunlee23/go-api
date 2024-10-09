package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProofRepository interface {
	Create(proof *models.Proof) error
	GetByID(id uint) (*models.Proof, error)
	GetAll(ctx *gin.Context) ([]models.Proof, error)
	Update(proof *models.Proof) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
