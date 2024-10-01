package routes

import (
	Handlers2 "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Handler  interface{}
	Register func(router *gin.Engine, Handler interface{})
}

func InitRoutes(productHandler *Handlers2.ProductHandler, productPartHandler *Handlers2.ProductPartHandler, denyRuleHandler *Handlers2.DenyRuleHandler, attributeHandler *Handlers2.AttributeHandler) []Route {
	return []Route{
		{Handler: productHandler, Register: func(r *gin.Engine, c interface{}) {
			ProductRoutes(r, c.(*Handlers2.ProductHandler))
		}},
		{Handler: productPartHandler, Register: func(r *gin.Engine, c interface{}) {
			ProductPartRoutes(r, c.(*Handlers2.ProductPartHandler))
		}},
		{Handler: denyRuleHandler, Register: func(r *gin.Engine, c interface{}) {
			DenyRuleRoutes(r, c.(*Handlers2.DenyRuleHandler))
		}},
		{Handler: attributeHandler, Register: func(r *gin.Engine, c interface{}) {
			AttributeRoutes(r, c.(*Handlers2.AttributeHandler))
		}},
	}
}

func RegisterRoutes(router *gin.Engine, routes []Route) {
	for _, route := range routes {
		route.Register(router, route.Handler)
	}
}
