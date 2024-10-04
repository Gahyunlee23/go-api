package services

import (
	"main-admin-api/internal/models"

	"github.com/gin-gonic/gin"
)

type AttributeCategoryService interface {
	GetAllAttributesCategories(ctx *gin.Context) (*models.ListResponse[models.AttributeCategory], error)
	GetAttributesCategoryByID(id uint, ctx *gin.Context) (*models.DetailResponse[models.AttributeCategory, models.Attribute], error)
}
