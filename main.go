package main

import (
	"log"
	_ "main-admin-api/docs"
	Handlers "main-admin-api/internal/api/handlers"
	"main-admin-api/internal/api/middleware"
	"main-admin-api/internal/api/routes"
	"main-admin-api/internal/services"
	"main-admin-api/pkg/config"
	"main-admin-api/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Database connection with config
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	serviceFactory := services.NewServiceFactory(db)
	productHandler := Handlers.NewProductHandler(serviceFactory.CreateProductService())
	productPartHandler := Handlers.NewProductPartHandler(serviceFactory.CreateProductPartService())
	denyRuleHandler := Handlers.NewDenyRuleHandler(serviceFactory.CreateDenyRuleService())
	attributeHandler := Handlers.NewAttributeHandler(serviceFactory.CreateAttributeService())
	fixedPriceHandler := Handlers.NewFixedPriceHandler(serviceFactory.CreateFixedPriceService())
	selectionRuleHandler := Handlers.NewSelectionRuleHandler(serviceFactory.CreateSelectionRuleService())
	attributeCategoryHandler := Handlers.NewAttributeCategory(serviceFactory.CreateAttributeCategoryService())
	productionTimeHandler := Handlers.NewProductionTimeHandler(serviceFactory.CreateProductionTimeService())
	proofHandler := Handlers.NewProofHandler(serviceFactory.CreateProofService())
	fileTypeHandler := Handlers.NewFileTypeHandler(serviceFactory.CreateFileTypeService())
	fileInspectionHandler := Handlers.NewFileInspectionHandler(serviceFactory.CreateFileInspectionService())

	router := gin.Default()
	router.RedirectTrailingSlash = false

	// convey config to CORS Middleware
	router.Use(middleware.SetupCORS(cfg))

	allRoutes := routes.InitRoutes(productHandler, productPartHandler, denyRuleHandler, attributeHandler,
		fixedPriceHandler, selectionRuleHandler, attributeCategoryHandler, productionTimeHandler,
		proofHandler, fileTypeHandler, fileInspectionHandler)

	routes.RegisterRoutes(router, allRoutes)
	routes.SwaggerRoutes(router)

	if err := router.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
