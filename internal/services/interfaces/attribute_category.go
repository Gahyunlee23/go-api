package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type AttributeCategoryService interface {
	GetAllAttributesCategories(ctx *gin.Context) (*models.ListResponse[models.AttributeCategory], error)
	GetAttributesCategoryByID(id uint, ctx *gin.Context) (*models.AttributeCategory, error)
	CreateAttributeCategory(attributeCategory *models.AttributeCategory) error
	UpdateAttributeCategory(urlID uint, attributeCategory *models.AttributeCategory) error
	ArchiveAttributeCategory(id uint) error
}
