package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main-admin-api/controllers"
	"main-admin-api/models"
	"main-admin-api/repository"
	"main-admin-api/routes"
	"main-admin-api/services"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:63383)/product_configurator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB: %v", err)
	}

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := services.NewProductServiceImpl(productRepository)
	productController := controllers.NewProductController(productService)

	router := gin.Default()

	routes.ProductRoutes(router, productController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
