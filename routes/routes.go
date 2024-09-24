package routes

import (
	"main-admin-api/controllers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Controller interface{}
	Register   func(router *gin.Engine, controller interface{})
}

func InitRoutes(productController *controllers.ProductController, productPartController *controllers.ProductPartController, denyRuleController *controllers.DenyRuleController) []Route {
	return []Route{
		{Controller: productController, Register: func(r *gin.Engine, c interface{}) {
			ProductRoutes(r, c.(*controllers.ProductController))
		}},
		{Controller: productPartController, Register: func(r *gin.Engine, c interface{}) {
			ProductPartRoutes(r, c.(*controllers.ProductPartController))
		}},
		{Controller: denyRuleController, Register: func(r *gin.Engine, c interface{}) {
			DenyRuleRoutes(r, c.(*controllers.DenyRuleController))
		}},
	}
}

func RegisterRoutes(router *gin.Engine, routes []Route) {
	for _, route := range routes {
		route.Register(router, route.Controller)
	}
}
