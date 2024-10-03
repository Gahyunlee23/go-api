package routes

import (
	handler "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SelectionRoutes(router *gin.Engine, handler *handler.SelectionRuleHandler) {
	SelectionRuleGroup := router.Group("/selection-rules")
	{
		SelectionRuleGroup.GET("/", handler.GetAllSelectionRules)
		SelectionRuleGroup.GET("/:id", handler.GetSelectionRuleByID)
		SelectionRuleGroup.POST("/", handler.CreateSelectionRule)
		SelectionRuleGroup.PUT("/:id", handler.UpdateSelectionRule)
		SelectionRuleGroup.DELETE("/:id", handler.DeleteSelectionRule)
	}
}
