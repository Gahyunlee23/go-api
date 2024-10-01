package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type AttributeService interface {
	CreateAttribute(attribute *models.Attribute, ctx *gin.Context) error
	GetAttributeByID(id uint) (*models.Attribute, error)
	GetAllAttributes(ctx *gin.Context) ([]models.Attribute, error)
	UpdateAttribute(attribute *models.Attribute, ctx *gin.Context) error
	DeleteAttribute(id uint) error
	ArchiveAttribute(id uint) error
}
