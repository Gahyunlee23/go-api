package routes

import (
	"main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func DenyRuleRoutes(router *gin.Engine, handler *handler.DenyRuleHandler) {
	DenyRuleGroup := router.Group("/deny-rules")
	{
		DenyRuleGroup.POST("/", handler.CreateDenyRule)
		DenyRuleGroup.GET("/", handler.GetAllDenyRules)
		DenyRuleGroup.GET("/:id", handler.GetDenyRuleByID)
		DenyRuleGroup.PUT("/:id", handler.UpdateDenyRule)
		DenyRuleGroup.DELETE("/:id", handler.DeleteDenyRule)
	}
}
