package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type FileInspectionRepository interface {
	Create(FileInspection *models.FileInspection) error
	GetByID(id uint) (*models.FileInspection, error)
	GetAll(ctx *gin.Context) ([]models.FileInspection, error)
	Update(FileInspection *models.FileInspection) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
