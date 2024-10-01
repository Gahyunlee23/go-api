package routes

import (
	Handlers "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Handler  interface{}
	Register func(router *gin.Engine, Handler interface{})
}

func InitRoutes(productHandler *Handlers.ProductHandler, productPartHandler *Handlers.ProductPartHandler, denyRuleHandler *Handlers.DenyRuleHandler, attributeHandler *Handlers.AttributeHandler, fixedPriceHandler *Handlers.FixedPriceHandler) []Route {
	return []Route{
		{Handler: productHandler, Register: func(r *gin.Engine, c interface{}) {
			ProductRoutes(r, c.(*Handlers.ProductHandler))
		}},
		{Handler: productPartHandler, Register: func(r *gin.Engine, c interface{}) {
			ProductPartRoutes(r, c.(*Handlers.ProductPartHandler))
		}},
		{Handler: denyRuleHandler, Register: func(r *gin.Engine, c interface{}) {
			DenyRuleRoutes(r, c.(*Handlers.DenyRuleHandler))
		}},
		{Handler: attributeHandler, Register: func(r *gin.Engine, c interface{}) {
			AttributeRoutes(r, c.(*Handlers.AttributeHandler))
		}},
		{Handler: fixedPriceHandler, Register: func(r *gin.Engine, c interface{}) {
			FixedPriceRoutes(r, c.(*Handlers.FixedPriceHandler))
		}},
	}
}

func RegisterRoutes(router *gin.Engine, routes []Route) {
	for _, route := range routes {
		route.Register(router, route.Handler)
	}
}
