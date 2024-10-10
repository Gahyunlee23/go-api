package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type FileTypeService interface {
	CreateFileType(fileType *models.FileType) error
	GetFileTypeByID(id uint) (*models.FileType, error)
	GetAllFileTypes(ctx *gin.Context) (*models.ListResponse[models.FileType], error)
	UpdateFileType(urlID uint, fileType *models.FileType) error
	ArchiveFileType(id uint) error
}
