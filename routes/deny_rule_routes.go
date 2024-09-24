package routes

import (
	"main-admin-api/controllers"

	"github.com/gin-gonic/gin"
)

func DenyRuleRoutes(router *gin.Engine, controller *controllers.DenyRuleController) {
	DenyRuleGroup := router.Group("/deny-rules")
	{
		DenyRuleGroup.POST("/", controller.CreateDenyRule)
		DenyRuleGroup.GET("/", controller.GetAllDenyRules)
		DenyRuleGroup.GET("/:id", controller.GetDenyRuleByID)
		DenyRuleGroup.PUT("/:id", controller.UpdateDenyRule)
		DenyRuleGroup.DELETE("/:id", controller.DeleteDenyRule)
	}
}
