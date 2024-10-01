package main

import (
	"log"
	_ "main-admin-api/docs"
	controllers2 "main-admin-api/internal/api/handlers"
	"main-admin-api/internal/api/middleware"
	"main-admin-api/internal/api/routes"
	"main-admin-api/internal/services"
	"main-admin-api/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	serviceFactory := services.NewServiceFactory(db)
	productController := controllers2.NewProductController(serviceFactory.CreateProductService())
	productPartController := controllers2.NewProductPartController(serviceFactory.CreateProductPartService())
	denyRuleController := controllers2.NewDenyRuleController(serviceFactory.CreateDenyRuleService())
	attributeController := controllers2.NewAttributeController(serviceFactory.CreateAttributeService())

	router := gin.Default()
	router.RedirectTrailingSlash = false
	router.Use(middleware.SetupCORS())

	allRoutes := routes.InitRoutes(productController, productPartController, denyRuleController, attributeController)

	routes.RegisterRoutes(router, allRoutes)

	routes.SwaggerRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
