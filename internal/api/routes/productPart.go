package routes

import (
	"main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func ProductPartRoutes(router *gin.Engine, controller *controllers.ProductPartController) {
	productPartGroup := router.Group("/product-parts")
	{
		productPartGroup.POST("/", controller.CreateProductPart)
		productPartGroup.GET("/", controller.GetAllProductParts)
		productPartGroup.GET("/:id", controller.GetProductPartByID)
		productPartGroup.PUT("/:id", controller.UpdateProductPart)
		productPartGroup.DELETE("/:id", controller.DeleteProductPart)
	}
}
