package handler

import (
	_ "log"
	"main-admin-api/internal/api/customerrors"
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
// @Tags Product
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
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.productService.CreateProduct(&product, ctx); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags Product
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [get]
func (c *ProductHandler) GetProductByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Product ID"})
		return
	}

	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use the 'code', 'id', 'name', or 'type' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', 'name', and 'type' parameters for cross-field AND search.
// @Description - sort[any_field]=asc or sort[ant_field]=desc
// @Description Example: /?search=keyword&code=abc&name=test&sort[name]=asc&sort[id]=desc
// @Tags Product
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Param type query string false "Filter by type field (partial match)"
// @Produce  json
// @Success 200 {array} models.Product
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/ [get]
func (c *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update the details of an existing product by providing the updated JSON payload
// @Tags Product
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Param   product  body  models.Product  true  "Updated product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [put]
func (c *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Product ID"})
		return
	}

	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.productService.UpdateProduct(uint(id), &product, ctx); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a single product by its ID
// @Tags Product
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [delete]
func (c *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Product ID"})
		return
	}

	if err := c.productService.ArchiveProduct(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "id": id})
}
