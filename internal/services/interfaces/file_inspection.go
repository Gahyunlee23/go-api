package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type FileInspectionService interface {
	CreateFileInspection(fileInspection *models.FileInspection) error
	GetFileInspectionByID(id uint) (*models.FileInspection, error)
	GetAllFileInspections(ctx *gin.Context) (*models.ListResponse[models.FileInspection], error)
	UpdateFileInspection(urlID uint, fileInspection *models.FileInspection) error
	ArchiveFileInspection(id uint) error
}
