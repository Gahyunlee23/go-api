package handler

import (
	"main-admin-api/internal/models"
	"main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttributeHandler struct {
	attributeService services.AttributeService
}

func NewAttributeHandler(service services.AttributeService) *AttributeHandler {
	return &AttributeHandler{attributeService: service}
}

// CreateAttribute godoc
// @Summary Create a new attribute
// @Description Create a attribute with the provided JSON payload
// @Tags Attributes
// @Accept  json
// @Produce  json
// @Param   Attribute  body  models.Attribute  true  "Attribute data"
// @Success 200 {object} models.Attribute
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/ [post]
func (c *AttributeHandler) CreateAttribute(ctx *gin.Context) {
	var attribute models.Attribute

	if err := ctx.ShouldBindJSON(&attribute); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.attributeService.CreateAttribute(&attribute, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"attribute": attribute})
}

// GetAttributeByID godoc
// @Summary Get Attribute by ID
// @Description Get a single attribute by its ID
// @Tags Attributes
// @Produce  json
// @Param   id  path  int  true  "Attribute ID"
// @Success 200 {object} models.Attribute
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Attribute not found"
// @Router /attributes/{id} [get]
func (c *AttributeHandler) GetAttributeByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attribute, err := c.attributeService.GetAttributeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, attribute)
}

// GetAllAttributes godoc
// @Summary Get all attribute
// @Description Retrieve a list of all attributes
// @Tags Attributes
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search term for filtering by name or code"
// @Produce  json
// @Success 200 {array} models.Attribute
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/ [get]
func (c *AttributeHandler) GetAllAttributes(ctx *gin.Context) {
	attributes, err := c.attributeService.GetAllAttributes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, attributes)
}

// UpdateAttribute godoc
// @Summary Update an existing attribute
// @Description Update the details of an existing attribute by providing the updated JSON payload
// @Tags Attributes
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Attribute ID"
// @Param   product  body  models.Attribute  true  "Updated Attribute data"
// @Success 200 {object} models.Attribute
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/{id} [put]
func (c *AttributeHandler) UpdateAttribute(ctx *gin.Context) {
	// Extract ID from the URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attribute ID"})
		return
	}

	// Bind JSON to attribute struct
	var attribute models.Attribute
	if err := ctx.ShouldBindJSON(&attribute); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service layer to update the attribute
	if err := c.attributeService.UpdateAttribute(uint(id), &attribute, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, attribute)
}

// DeleteAttribute godoc
// @Summary Delete an attribute by ID
// @Description Delete a single attribute by its ID
// @Tags Attributes
// @Produce json
// @Param id path int true "Attribute ID"
// @Success 200 {object} map[string]interface{} "Attribute deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/{id} [delete]
func (c *AttributeHandler) DeleteAttribute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid ID": err.Error()})
		return
	}
	if err := c.attributeService.ArchiveAttribute(uint(id)); err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Attribute deleted successfully": id})
}
