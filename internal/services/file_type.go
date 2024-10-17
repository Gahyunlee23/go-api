package services

import (
	"fmt"
	"main-admin-api/internal/models"
	repository "main-admin-api/internal/repository/interfaces"
	services "main-admin-api/internal/services/interfaces"
	"main-admin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type fileTypeService struct {
	fileTypeRepository repository.FileTypeRepository
}

func NewFileTypeService(repository repository.FileTypeRepository) services.FileTypeService {
	return &fileTypeService{fileTypeRepository: repository}
}

func (s *fileTypeService) CreateFileType(fileType *models.FileType) error {
	return s.fileTypeRepository.Create(fileType)
}

func (s *fileTypeService) GetFileTypeByID(id uint) (*models.FileType, error) {
	return s.fileTypeRepository.GetByID(id)
}

func (s *fileTypeService) GetAllFileTypes(ctx *gin.Context) (*models.ListResponse[models.FileType], error) {
	fileTypes, err := s.fileTypeRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.fileTypeRepository.Count(ctx)
	if err != nil {
		return nil, err
	}

	response := models.NewListResponse(totalCount, fileTypes)
	return &response, nil

}

func (s *fileTypeService) UpdateFileType(urlID uint, fileType *models.FileType) error {
	// Verify that URL ID matches the attribute ID
	if err := utils.ValidateID(urlID, fileType.ID); err != nil {
		return fmt.Errorf("ID validation failed: %w", err)
	}

	if _, err := utils.ValidateAndFetchEntity(s.fileTypeRepository, urlID, "File Type"); err != nil {
		return fmt.Errorf("validate and fetch entity: %w", err)
	}

	if err := s.fileTypeRepository.Update(fileType); err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}

	return nil
}

func (s *fileTypeService) ArchiveFileType(id uint) error {
	if _, err := utils.ValidateAndFetchEntity(s.fileTypeRepository, id, "File Type"); err != nil {
		return fmt.Errorf("validate and fetch entity: %w", err)
	}
	if err := s.fileTypeRepository.Archive(id); err != nil {
		return fmt.Errorf("failed to archive entity: %w", err)
	}
	return nil
}
