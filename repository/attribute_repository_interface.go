package repository

import (
	"main-admin-api/models"

	"github.com/gin-gonic/gin"
)

type AttributeRepositoryInterface interface {
	Create(attribute *models.Attribute) error
	GetByID(ID uint) (*models.Attribute, error)
	GetAll(ctx *gin.Context) ([]models.Attribute, error)
	Update(attribute *models.Attribute) error
	Delete(ID uint) error
	Archive(ID uint) error
}
