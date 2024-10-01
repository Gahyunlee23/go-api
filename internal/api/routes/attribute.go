package routes

import (
	"main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func AttributeRoutes(router *gin.Engine, controller *controllers.AttributeController) {
	AttributeGroup := router.Group("/attributes")
	{
		AttributeGroup.POST("/", controller.CreateAttribute)
		AttributeGroup.GET("/", controller.GetAllAttributes)
		AttributeGroup.GET("/:id", controller.GetAttributeByID)
		AttributeGroup.PUT("/:id", controller.UpdateAttribute)
		AttributeGroup.DELETE("/:id", controller.DeleteAttribute)
	}
}
