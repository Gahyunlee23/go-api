package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FixedPriceHandler struct {
	FixedPriceService services.FixedPriceService
}

func NewFixedPriceHandler(service services.FixedPriceService) *FixedPriceHandler {
	return &FixedPriceHandler{FixedPriceService: service}
}

// CreateFixedPrice godoc
// @Summary Create a new Fixed Price
// @Description Create a Fixed Price with the provided JSON payload
// @Tags FixedPrice
// @Accept  json
// @Produce  json
// @Param   FixedPrice  body  models.FixedPrice  true  "Fixed Price data"
// @Success 200 {object} models.FixedPrice
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/ [post]
func (c *FixedPriceHandler) CreateFixedPrice(ctx *gin.Context) {
	var fixedPrice models.FixedPrice

	if err := ctx.ShouldBindJSON(&fixedPrice); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.FixedPriceService.CreateFixedPrice(ctx, &fixedPrice); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, fixedPrice)
}

// GetFixedPriceByID godoc
// @Summary Get Fixed Price by ID
// @Description Get a single Fixed Price by its ID
// @Tags FixedPrice
// @Produce  json
// @Param   id  path  int  true  "Fixed Price ID"
// @Success 200 {object} models.FixedPrice
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/{id} [get]
func (c *FixedPriceHandler) GetFixedPriceByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Fixed Price ID"})
		return
	}

	fixedPrice, err := c.FixedPriceService.GetFixedPriceByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fixedPrice)
}

// GetAllFixedPrices godoc
// @Summary Get all FixedPrices
// @Description Retrieve a list of all FixedPrices.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use the 'code', 'id', or 'name' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', and 'name' parameters for cross-field AND search.
// @Description - sort[any_field]=asc or sort[ant_field]=desc
// @Description Example: /?search=keyword&code=abc&name=test&sort[name]=asc&sort[id]=desc
// @Tags FixedPrice
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Produce  json
// @Success 200 {array} models.FixedPrice
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/ [get]
func (c *FixedPriceHandler) GetAllFixedPrices(ctx *gin.Context) {
	fixedPrices, err := c.FixedPriceService.GetAllFixedPrices(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fixedPrices)
}

// UpdateFixedPrice godoc
// @Summary Update an existing Fixed Price
// @Description Update the details of an existing FixedPrice by providing the updated JSON payload
// @Tags FixedPrice
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "FixedPrice ID"
// @Param   product  body  models.FixedPrice  true  "Updated Fixed Price data"
// @Success 200 {object} models.FixedPrice
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/{id} [put]
func (c *FixedPriceHandler) UpdateFixedPrice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Fixed Price ID"})
		return
	}

	var fixedPrice models.FixedPrice
	if err := ctx.ShouldBindJSON(&fixedPrice); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.FixedPriceService.UpdateFixedPrice(uint(id), &fixedPrice, ctx); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fixedPrice)
}

// DeleteFixedPrice godoc
// @Summary Delete an Fixed Price by ID
// @Description Delete a single Fixed Price by its ID
// @Tags FixedPrice
// @Produce json
// @Param id path int true "Fixed Price ID"
// @Success 200 {object} map[string]interface{} "Fixed Price deleted successfully"
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/{id} [delete]
func (c *FixedPriceHandler) DeleteFixedPrice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Fixed Price ID"})
		return
	}

	if err := c.FixedPriceService.ArchiveFixedPrice(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Fixed Price deleted successfully", "id": id})
}
