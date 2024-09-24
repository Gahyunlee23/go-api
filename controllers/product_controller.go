package controllers

import (
	"log"
	"main-admin-api/models"
	"main-admin-api/services"
	"main-admin-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductServiceInterface
}

func NewProductController(service services.ProductServiceInterface) *ProductController {
	return &ProductController{productService: service}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a product with the provided JSON payload
// @Tags products
// @Accept  json
// @Produce  json
// @Param   product  body  models.Product  true  "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/ [post]
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error parsing JSON:", err)
		return
	}

	var err error
	product.RenamingRules, err = utils.MarshalAndAssignJSON(product.RenamingRules, "renaming_rules", ctx)
	if err != nil {
		return
	}

	product.OrderRules, err = utils.MarshalAndAssignJSON(product.OrderRules, "order_rules", ctx)
	if err != nil {
		return
	}

	product.QuantitiesSelection, err = utils.MarshalAndAssignJSON(product.QuantitiesSelection, "quantities_selection", ctx)
	if err != nil {
		return
	}

	if err := c.productService.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"product": product})
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags products
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /products/{id} [get]
func (c *ProductController) GetProductByID(ctx *gin.Context) {
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
// @Tags products
// @Produce  json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/ [get]
func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update the details of an existing product by providing the updated JSON payload
// @Tags products
// @Accept  json
// @Produce  json
// @Param   product  body  models.Product  true  "Updated product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [put]
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	product.RenamingRules, err = utils.MarshalAndAssignJSON(product.RenamingRules, "renaming_rules", ctx)
	if err != nil {
		return
	}

	product.OrderRules, err = utils.MarshalAndAssignJSON(product.OrderRules, "order_rules", ctx)
	if err != nil {
		return
	}

	product.QuantitiesSelection, err = utils.MarshalAndAssignJSON(product.QuantitiesSelection, "quantities_selection", ctx)
	if err != nil {
		return
	}

	if err := c.productService.UpdateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a single product by its ID
// @Tags products
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [delete]
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.productService.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
