package services

import (
	"errors"
	"fmt"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type proofService struct {
	proofRepository repository.ProofRepository
}

func NewProofService(repository repository.ProofRepository) services.ProofService {
	return &proofService{proofRepository: repository}
}

func (s *proofService) CreateProof(proof *models.Proof) error {
	return s.proofRepository.Create(proof)
}

func (s *proofService) GetProofByID(id uint) (*models.Proof, error) {
	return s.proofRepository.GetByID(id)
}

func (s *proofService) GetAllProofs(ctx *gin.Context) (*models.ListResponse[models.Proof], error) {
	proofs, err := s.proofRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	totalCount, err := s.proofRepository.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	response := models.NewListResponse(totalCount, proofs)
	return &response, nil
}

func (s *proofService) UpdateProof(urlID uint, proof *models.Proof) error {
	if urlID != proof.ID {
		return errors.New("proof ID in URL does not match the ID in the request body")
	}

	if _, err := utils.ValidateAndFetchEntity(s.proofRepository, urlID, "Proof"); err != nil {
		return fmt.Errorf("failed to validate proof: %w", err)
	}

	if err := s.proofRepository.Update(proof); err != nil {
		return fmt.Errorf("failed to update proof: %w", err)
	}

	return nil
}

func (s *proofService) ArchiveProof(id uint) error {
	if _, err := utils.ValidateAndFetchEntity(s.proofRepository, id, "Proof"); err != nil {
		return fmt.Errorf("failed to validate proof: %w", err)
	}

	if err := s.proofRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to delete proof: %w", err)
	}

	return nil
}
