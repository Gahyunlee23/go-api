package routes

import (
	controllers2 "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Controller interface{}
	Register   func(router *gin.Engine, controller interface{})
}

func InitRoutes(productController *controllers2.ProductController, productPartController *controllers2.ProductPartController, denyRuleController *controllers2.DenyRuleController, attributeController *controllers2.AttributeController) []Route {
	return []Route{
		{Controller: productController, Register: func(r *gin.Engine, c interface{}) {
			ProductRoutes(r, c.(*controllers2.ProductController))
		}},
		{Controller: productPartController, Register: func(r *gin.Engine, c interface{}) {
			ProductPartRoutes(r, c.(*controllers2.ProductPartController))
		}},
		{Controller: denyRuleController, Register: func(r *gin.Engine, c interface{}) {
			DenyRuleRoutes(r, c.(*controllers2.DenyRuleController))
		}},
		{Controller: attributeController, Register: func(r *gin.Engine, c interface{}) {
			AttributeRoutes(r, c.(*controllers2.AttributeController))
		}},
	}
}

func RegisterRoutes(router *gin.Engine, routes []Route) {
	for _, route := range routes {
		route.Register(router, route.Controller)
	}
}
