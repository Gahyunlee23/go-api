package handler

import (
	"log"
	"main-admin-api/internal/models"
	"main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{productService: service}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a product with the provided JSON payload
// @Tags Products
// @Accept  json
// @Produce  json
// @Param   product  body  models.Product  true  "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/ [post]
func (c *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error parsing JSON:", err)
		return
	}

	if err := c.productService.CreateProduct(&product, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"product": product})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags Products
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /products/{id} [get]
func (c *ProductHandler) GetProductByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags Products
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search term for filtering by name or code"
// @Produce  json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/ [get]
func (c *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update the details of an existing product by providing the updated JSON payload
// @Tags Products
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Param   product  body  models.Product  true  "Updated product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [put]
func (c *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.productService.UpdateProduct(uint(id), &product, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a single product by its ID
// @Tags Products
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [delete]
func (c *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid ID": err.Error()})
		return
	}
	if err := c.productService.ArchiveProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Product deleted successfully": id})
}
