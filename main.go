package main

import (
	"log"
	"main-admin-api/controllers"
	"main-admin-api/database"
	_ "main-admin-api/docs"
	"main-admin-api/middleware"
	"main-admin-api/routes"
	"main-admin-api/services"

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
	productController := controllers.NewProductController(serviceFactory.CreateProductService())
	productPartController := controllers.NewProductPartController(serviceFactory.CreateProductPartService())
	denyRuleController := controllers.NewDenyRuleController(serviceFactory.CreateDenyRuleService())
	attributeController := controllers.NewAttributeController(serviceFactory.CreateAttributeService())

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
