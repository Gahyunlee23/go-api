package routes

import (
	Handlers "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func ProductionTimeRoutes(router *gin.Engine, handler *Handlers.ProductionTimeHandler) {
	productionTimeRoute := router.Group("/production-times")
	{
		productionTimeRoute.GET("/", handler.GetAllProductionTime)
		productionTimeRoute.POST("/", handler.CreateProductionTime)
		productionTimeRoute.GET("/:id", handler.GetProductionTimeByID)
		productionTimeRoute.PUT("/:id", handler.UpdateProductionTime)
		productionTimeRoute.DELETE("/:id", handler.DeleteProductionTime)
	}
}
