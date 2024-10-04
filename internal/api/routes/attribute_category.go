package routes

import (
	handler "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func AttributeCategoryRoutes(router *gin.Engine, handler *handler.AttributeCategoryHandler) {
	AttributeCategory := router.Group("/attribute-categories")
	{
		AttributeCategory.GET("/", handler.GetAllAttributeCategory)
		AttributeCategory.GET("/:id", handler.GetAttributeCategoryByID)
	}
}
