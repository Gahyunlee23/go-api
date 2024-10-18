package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductionTimeHandler struct {
	productionTimeService services.ProductionTimeService
}

func NewProductionTimeHandler(service services.ProductionTimeService) *ProductionTimeHandler {
	return &ProductionTimeHandler{productionTimeService: service}
}

// CreateProductionTime godoc
// @Summary Create a new productionTime
// @Tags ProductionTime
// @Accept json
// @Produce  json
// @Param   productionTime  body  models.ProductionTime  true  "Production Time data"
// @Success 200 {object} models.ProductionTime
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /production-times/ [post]
func (c *ProductionTimeHandler) CreateProductionTime(ctx *gin.Context) {
	var productionTime models.ProductionTime

	if err := ctx.ShouldBindJSON(&productionTime); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.productionTimeService.CreateProductionTime(&productionTime); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, productionTime)
}

// GetProductionTimeByID godoc
// @Summary Get Production Time by ID
// @Description Get a single Production Time by its ID
// @Tags ProductionTime
// @Produce  json
// @Param   id  path  int  true  "Production Time ID"
// @Success 200 {object} models.ProductionTime
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /production-times/{id} [get]
func (c *ProductionTimeHandler) GetProductionTimeByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Product ionTime ID"})
		return
	}

	productionTime, err := c.productionTimeService.GetProductionTimeByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, productionTime)
}

// GetAllProductionTime godoc
// @Summary Get all productionTime
// @Description Retrieve a list of all productionTime.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use the 'code', 'id', 'name', or 'time' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', 'name', and 'time' parameters for cross-field AND search.
// @Description Example: /productionTime?search=keyword&code=abc&name=test
// @Tags ProductionTime
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Param content_type query string false "Filter by time field (partial match)"
// @Produce  json
// @Success 200 {array} models.ProductionTime
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /production-times/ [get]
func (c *ProductionTimeHandler) GetAllProductionTime(ctx *gin.Context) {
	productionTime, err := c.productionTimeService.GetAllProductionTimes(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, productionTime)
}

// UpdateProductionTime godoc
// @Summary Update an existing production Time
// @Description Update the details of an existing production Time by providing the updated JSON payload
// @Tags ProductionTime
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Production Time ID"
// @Param   product  body  models.ProductionTime  true  "Updated product data"
// @Success 200 {object} models.ProductionTime "successfully updated production time"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 400 {object} models.ErrorResponse "ID mismatch error"
// @Failure 400 {object} models.ErrorResponse "Invalid ID in request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /production-times/{id} [put]
func (c *ProductionTimeHandler) UpdateProductionTime(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Production Time ID"})
		return
	}

	var productionTime models.ProductionTime
	if err := ctx.ShouldBindJSON(&productionTime); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.productionTimeService.UpdateProductionTime(uint(id), &productionTime); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, productionTime)
}

// DeleteProductionTime godoc
// @Summary Delete a product ionTime by ID
// @Description Delete a single production Time by its ID
// @Tags ProductionTime
// @Produce  json
// @Param   id  path  int  true  "Production Time ID"
// @Success 200 {object} map[string]interface{} "Production Time deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /production-times/{id} [delete]
func (c *ProductionTimeHandler) DeleteProductionTime(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Production Time ID"})
		return
	}

	if err := c.productionTimeService.ArchiveProductionTime(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Production Time deleted successfully", "id": id})
}
