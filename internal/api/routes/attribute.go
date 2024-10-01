package routes

import (
	"main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func AttributeRoutes(router *gin.Engine, handler *handler.AttributeHandler) {
	AttributeGroup := router.Group("/attributes")
	{
		AttributeGroup.POST("/", handler.CreateAttribute)
		AttributeGroup.GET("/", handler.GetAllAttributes)
		AttributeGroup.GET("/:id", handler.GetAttributeByID)
		AttributeGroup.PUT("/:id", handler.UpdateAttribute)
		AttributeGroup.DELETE("/:id", handler.DeleteAttribute)
	}
}
