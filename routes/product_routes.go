package routes

import (
	"github.com/gin-gonic/gin"
	"main-admin-api/controllers"
)

func ProductRoutes(router *gin.Engine, controller *controllers.ProductController) {
	productGroup := router.Group("/products")
	{
		productGroup.POST("/", controller.CreateProduct)
		productGroup.GET("/", controller.GetAllProducts)
		productGroup.GET("/:id", controller.GetProductByID)
		productGroup.PUT("/", controller.UpdateProduct)
		productGroup.DELETE("/:id", controller.DeleteProduct)
	}
}
