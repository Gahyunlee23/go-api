package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type ProofService interface {
	CreateProof(proof *models.Proof) error
	GetProofByID(id uint) (*models.Proof, error)
	GetAllProofs(ctx *gin.Context) (*models.ListResponse[models.Proof], error)
	UpdateProof(urlID uint, proof *models.Proof) error
	ArchiveProof(id uint) error
}
