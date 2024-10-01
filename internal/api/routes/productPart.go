package routes

import (
	"main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func ProductPartRoutes(router *gin.Engine, handler *handler.ProductPartHandler) {
	productPartGroup := router.Group("/product-parts")
	{
		productPartGroup.POST("/", handler.CreateProductPart)
		productPartGroup.GET("/", handler.GetAllProductParts)
		productPartGroup.GET("/:id", handler.GetProductPartByID)
		productPartGroup.PUT("/:id", handler.UpdateProductPart)
		productPartGroup.DELETE("/:id", handler.DeleteProductPart)
	}
}
