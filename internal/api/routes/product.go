package routes

import (
	"main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, handler *handler.ProductHandler) {
	productGroup := router.Group("/products")
	{
		productGroup.POST("/", handler.CreateProduct)
		productGroup.GET("/", handler.GetAllProducts)
		productGroup.GET("/:id", handler.GetProductByID)
		productGroup.PUT("/:id", handler.UpdateProduct)
		productGroup.DELETE("/:id", handler.DeleteProduct)
	}
}
