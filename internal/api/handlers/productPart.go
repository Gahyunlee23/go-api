package handler

import (
	_ "encoding/json"
	"main-admin-api/internal/api/errors"
	"main-admin-api/internal/models"
	"main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "gorm.io/datatypes"
)

type ProductPartHandler struct {
	productPartService services.ProductPartService
}

func NewProductPartHandler(service services.ProductPartService) *ProductPartHandler {
	return &ProductPartHandler{productPartService: service}
}

// CreateProductPart godoc
// @Summary Create a new product part
// @Tags ProductPart
// @Accept json
// @Produce  json
// @Param   productPart  body  models.ProductPart  true  "Product Part data"
// @Success 200 {object} models.ProductPart
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /product-parts/ [post]
func (c *ProductPartHandler) CreateProductPart(ctx *gin.Context) {
	var productPart models.ProductPart

	if err := ctx.ShouldBindJSON(&productPart); err != nil {
		errors.HandleError(ctx, &errors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.productPartService.CreateProductPart(&productPart, ctx); err != nil {
		errors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, productPart)
}

// GetProductPartByID godoc
// @Summary Get Product Part by ID
// @Description Get a single Product Part by its ID
// @Tags ProductPart
// @Produce  json
// @Param   id  path  int  true  "Product Part ID"
// @Success 200 {object} models.ProductPart
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Product Part not found"
// @Router /product-parts/{id} [get]
func (c *ProductPartHandler) GetProductPartByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		errors.HandleError(ctx, &errors.ValidationError{Field: "id", Message: "Invalid Product Part ID"})
		return
	}

	productPart, err := c.productPartService.GetProductPartByID(uint(id))
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, productPart)
}

// GetAllProductParts godoc
// @Summary Get all product parts
// @Description Retrieve a list of all product parts
// @Tags ProductPart
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search term for filtering by name or code"
// @Produce  json
// @Success 200 {array} models.ProductPart
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /product-parts/ [get]
func (c *ProductPartHandler) GetAllProductParts(ctx *gin.Context) {
	productParts, err := c.productPartService.GetAllProductPart(ctx)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, productParts)
}

// UpdateProductPart godoc
// @Summary Update an existing product part
// @Description Update the details of an existing product part by providing the updated JSON payload
// @Tags ProductPart
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Product Part ID"
// @Param   product  body  models.ProductPart  true  "Updated product data"
// @Success 200 {object} models.ProductPart
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /product-parts/{id} [put]
func (c *ProductPartHandler) UpdateProductPart(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		errors.HandleError(ctx, &errors.ValidationError{Field: "id", Message: "Invalid Product Part ID"})
		return
	}

	var productPart models.ProductPart
	if err := ctx.ShouldBindJSON(&productPart); err != nil {
		errors.HandleError(ctx, &errors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.productPartService.UpdateProductPart(uint(id), &productPart, ctx); err != nil {
		errors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, productPart)
}

// DeleteProductPart godoc
// @Summary Delete a product part by ID
// @Description Delete a single product part by its ID
// @Tags ProductPart
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Success 200 {object} map[string]interface{} "Product Part deleted successfully"
// @Router /product-parts/{id} [delete]
func (c *ProductPartHandler) DeleteProductPart(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		errors.HandleError(ctx, &errors.ValidationError{Field: "id", Message: "Invalid Product Part ID"})
		return
	}

	if err := c.productPartService.ArchiveProductPart(uint(id)); err != nil {
		errors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product Part deleted successfully", "id": id})
}
