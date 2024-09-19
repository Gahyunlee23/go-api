package main

import (
	"log"
	"main-admin-api/controllers"
	_ "main-admin-api/docs"
	"main-admin-api/models"
	"main-admin-api/repository"
	"main-admin-api/routes"
	"main-admin-api/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowOrigins:     []string{"http://localhost:9001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.ProductRoutes(router, productController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
