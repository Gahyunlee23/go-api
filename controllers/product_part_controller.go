package controllers

import (
	_ "encoding/json"
	"log"
	"main-admin-api/models"
	"main-admin-api/services"
	"main-admin-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "gorm.io/datatypes"
)

type ProductPartController struct {
	productPartService services.ProductPartServiceInterface
}

func NewProductPartController(service services.ProductPartServiceInterface) *ProductPartController {
	return &ProductPartController{productPartService: service}
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
func (c *ProductPartController) CreateProductPart(ctx *gin.Context) {
	var productPart models.ProductPart

	if err := ctx.ShouldBindJSON(&productPart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error parsing JSON:", err)
		return
	}

	var err error
	productPart.Paper, err = utils.MarshalAndAssignJSON(productPart.Paper, "paper", ctx)
	if err != nil {
		return
	}

	productPart.Format, err = utils.MarshalAndAssignJSON(productPart.Format, "format", ctx)
	if err != nil {
		return
	}

	productPart.Pages, err = utils.MarshalAndAssignJSON(productPart.Pages, "pages", ctx)
	if err != nil {
		return
	}

	productPart.Colors, err = utils.MarshalAndAssignJSON(productPart.Colors, "colors", ctx)
	if err != nil {
		return
	}

	productPart.BookBinding, err = utils.MarshalAndAssignJSON(productPart.BookBinding, "bookBinding", ctx)
	if err != nil {
		return
	}

	productPart.Refinement, err = utils.MarshalAndAssignJSON(productPart.Refinement, "refinement", ctx)
	if err != nil {
		return
	}

	productPart.Finishing, err = utils.MarshalAndAssignJSON(productPart.Finishing, "finishing", ctx)
	if err != nil {
		return
	}

	productPart.DefaultSelections, err = utils.MarshalAndAssignJSON(productPart.DefaultSelections, "defaultSelections", ctx)
	if err != nil {
		return
	}

	if err := c.productPartService.CreateProductPart(&productPart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"productPart": productPart})
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
func (c *ProductPartController) GetProductPartByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid ID": err.Error()})
	}
	productPart, err := c.productPartService.GetProductPartByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product Part not found"})
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
func (c *ProductPartController) GetAllProductParts(ctx *gin.Context) {
	productParts, err := c.productPartService.GetAllProductPart(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"productParts": productParts})
}

// UpdateProductPart godoc
// @Summary Update an existing product part
// @Description Update the details of an existing product part by providing the updated JSON payload
// @Tags ProductPart
// @Accept  json
// @Produce  json
// @Param   product  body  models.ProductPart  true  "Updated product data"
// @Success 200 {object} models.ProductPart
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /product-parts/{id} [put]
func (c *ProductPartController) UpdateProductPart(ctx *gin.Context) {
	var productPart models.ProductPart
	if err := ctx.ShouldBindJSON(&productPart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	productPart.Paper, err = utils.MarshalAndAssignJSON(productPart.Paper, "paper", ctx)
	if err != nil {
		return
	}

	productPart.Format, err = utils.MarshalAndAssignJSON(productPart.Format, "format", ctx)
	if err != nil {
		return
	}

	productPart.Pages, err = utils.MarshalAndAssignJSON(productPart.Pages, "pages", ctx)
	if err != nil {
		return
	}

	productPart.Colors, err = utils.MarshalAndAssignJSON(productPart.Colors, "colors", ctx)
	if err != nil {
		return
	}

	productPart.BookBinding, err = utils.MarshalAndAssignJSON(productPart.BookBinding, "bookBinding", ctx)
	if err != nil {
		return
	}

	productPart.Refinement, err = utils.MarshalAndAssignJSON(productPart.Refinement, "refinement", ctx)
	if err != nil {
		return
	}

	productPart.Finishing, err = utils.MarshalAndAssignJSON(productPart.Finishing, "finishing", ctx)
	if err != nil {
		return
	}

	productPart.DefaultSelections, err = utils.MarshalAndAssignJSON(productPart.DefaultSelections, "defaultSelections", ctx)
	if err != nil {
		return
	}

	if err := c.productPartService.UpdateProductPart(&productPart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"productPart": productPart})
}

// DeleteProductPart godoc
// @Summary Delete a product part by ID
// @Description Delete a single product part by its ID
// @Tags ProductPart
// @Produce  json
// @Param   id  path  int  true  "Product ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Router /product-parts/{id} [delete]
func (c *ProductPartController) DeleteProductPart(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid ID": err.Error()})
	}
	if err := c.productPartService.ArchiveProductPart(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"productPart": nil})
}
