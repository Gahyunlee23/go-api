package repository

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type AttributeRepository interface {
	Create(attribute *models.Attribute) error
	GetByID(ID uint) (*models.Attribute, error)
	GetAll(ctx *gin.Context) ([]models.Attribute, error)
	Update(attribute *models.Attribute) error
	Delete(ID uint) error
	Archive(ID uint) error
}
