package routes

import (
	"main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, controller *controllers.ProductController) {
	productGroup := router.Group("/products")
	{
		productGroup.POST("/", controller.CreateProduct)
		productGroup.GET("/", controller.GetAllProducts)
		productGroup.GET("/:id", controller.GetProductByID)
		productGroup.PUT("/:id", controller.UpdateProduct)
		productGroup.DELETE("/:id", controller.DeleteProduct)
	}
}
