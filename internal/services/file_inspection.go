package services

import (
	"fmt"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type fileInspectionService struct {
	fileInspectionRepository repository.FileInspectionRepository
}

func NewFileInspectionService(repository repository.FileInspectionRepository) services.FileInspectionService {
	return &fileInspectionService{fileInspectionRepository: repository}
}

func (s *fileInspectionService) CreateFileInspection(fileInspection *models.FileInspection) error {
	return s.fileInspectionRepository.Create(fileInspection)
}

func (s *fileInspectionService) GetFileInspectionByID(id uint) (*models.FileInspection, error) {
	return s.fileInspectionRepository.GetByID(id)
}

func (s *fileInspectionService) GetAllFileInspections(ctx *gin.Context) (*models.ListResponse[models.FileInspection], error) {
	fileInspections, err := s.fileInspectionRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.fileInspectionRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, fileInspections)
	return &response, nil

}

func (s *fileInspectionService) UpdateFileInspection(urlID uint, fileInspection *models.FileInspection) error {
	// Verify that URL ID matches the attribute ID
	if err := utils.ValidateID(urlID, fileInspection.ID); err != nil {
		return fmt.Errorf("ID validation failed: %w", err)
	}

	if _, err := utils.ValidateAndFetchEntity(s.fileInspectionRepository, urlID, "File Inspection"); err != nil {
		return fmt.Errorf("validate and fetch entity: %w", err)
	}

	if err := s.fileInspectionRepository.Update(fileInspection); err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}

	return nil
}

func (s *fileInspectionService) ArchiveFileInspection(id uint) error {
	if _, err := utils.ValidateAndFetchEntity(s.fileInspectionRepository, id, "File Inspection"); err != nil {
		return fmt.Errorf("validate and fetch entity: %w", err)
	}
	if err := s.fileInspectionRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to archive entity: %w", err)
	}
	return nil
}
