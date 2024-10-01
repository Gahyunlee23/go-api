package routes

import (
	handler "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func FixedPriceRoutes(router *gin.Engine, handler *handler.FixedPriceHandler) {
	FixedPriceGroup := router.Group("/fixed-prices")
	{
		FixedPriceGroup.GET("/", handler.GetAllFixedPrices)
		FixedPriceGroup.POST("/", handler.CreateFixedPrice)
		FixedPriceGroup.GET("/:id", handler.GetFixedPriceByID)
		FixedPriceGroup.PUT("/:id", handler.UpdateFixedPrice)
		FixedPriceGroup.DELETE("/:id", handler.DeleteFixedPrice)
	}
}
