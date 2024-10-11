package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type FileTypeRepository interface {
	Create(fileType *models.FileType) error
	GetByID(id uint) (*models.FileType, error)
	GetAll(ctx *gin.Context) ([]models.FileType, error)
	Update(fileType *models.FileType) error
	Archive(id uint) error
	Count(ctx *gin.Context) (int64, error)
}
