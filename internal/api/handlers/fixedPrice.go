package handler

import (
	"main-admin-api/internal/models"
	"main-admin-api/internal/services/interfaces"
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
// @Tags FixedPrices
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.FixedPriceService.CreateFixedPrice(ctx, &fixedPrice); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"FixedPrice": fixedPrice})
}

// GetFixedPriceByID godoc
// @Summary Get FixedPrice by ID
// @Description Get a single FixedPrice by its ID
// @Tags FixedPrices
// @Produce  json
// @Param   id  path  int  true  "FixedPrice ID"
// @Success 200 {object} models.FixedPrice
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Fixed Price not found"
// @Router /fixed-prices/{id} [get]
func (c *FixedPriceHandler) GetFixedPriceByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	FixedPrice, err := c.FixedPriceService.GetFixedPriceByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, FixedPrice)
}

// GetAllFixedPrices godoc
// @Summary Get all FixedPrice
// @Description Retrieve a list of all FixedPrices
// @Tags FixedPrices
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search term for filtering by name or code"
// @Produce  json
// @Success 200 {array} models.FixedPrice
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/ [get]
func (c *FixedPriceHandler) GetAllFixedPrices(ctx *gin.Context) {
	FixedPrices, err := c.FixedPriceService.GetAllFixedPrices(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, FixedPrices)
}

// UpdateFixedPrice godoc
// @Summary Update an existing FixedPrice
// @Description Update the details of an existing FixedPrice by providing the updated JSON payload
// @Tags FixedPrices
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "FixedPrice ID"
// @Param   product  body  models.FixedPrice  true  "Updated FixedPrice data"
// @Success 200 {object} models.FixedPrice
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/{id} [put]
func (c *FixedPriceHandler) UpdateFixedPrice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	var fixedPrice models.FixedPrice
	if err := ctx.ShouldBindJSON(&fixedPrice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.FixedPriceService.UpdateFixedPrice(uint(id), ctx, &fixedPrice); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update FixedPrice"})
		return
	}
	ctx.JSON(http.StatusOK, fixedPrice)
}

// DeleteFixedPrice godoc
// @Summary Delete an FixedPrice by ID
// @Description Delete a single FixedPrice by its ID
// @Tags FixedPrices
// @Produce json
// @Param id path int true "Fixed Price ID"
// @Success 200 {object} map[string]interface{} "Fixed Price deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /fixed-prices/{id} [delete]
func (c *FixedPriceHandler) DeleteFixedPrice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid ID": err.Error()})
		return
	}
	if err := c.FixedPriceService.ArchiveFixedPrice(uint(id)); err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Fixed Price deleted successfully": id})
}
