package integration

import (
	"errors"
	"fmt"
	"log"
	Handlers "main-admin-api/internal/api/handlers"
	"main-admin-api/internal/api/routes"
	"main-admin-api/internal/services"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testServer *http.Server

func connectDB() (*gorm.DB, error) {
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := "52548"
	dbname := "product_configurator"
	charset := os.Getenv("DB_CHARSET")
	parseTime := "utf8mb4"
	loc := "Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, password, host, port, dbname, charset, parseTime, loc)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupRouter() *gin.Engine {
	db, err := connectDB()
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

	allRoutes := routes.InitRoutes(
		productHandler, productPartHandler, denyRuleHandler,
		attributeHandler, fixedPriceHandler, selectionRuleHandler,
		attributeCategoryHandler, productionTimeHandler, proofHandler,
		fileTypeHandler, fileInspectionHandler,
	)

	routes.RegisterRoutes(router, allRoutes)

	return router
}

func TestMain(m *testing.M) {
	// Set to Test Mode
	gin.SetMode(gin.TestMode)

	currentDir, _ := os.Getwd()
	log.Printf("Current working directory: %s", currentDir)

	router := setupRouter()

	testServer = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := testServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("TestServer ListenAndServe: %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	code := m.Run()

	if err := testServer.Close(); err != nil {
		log.Printf("TestServer Close: %v", err)
	}

	os.Exit(code)
}
