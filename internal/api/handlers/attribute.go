package handler

import (
	"main-admin-api/internal/api/customerrors"
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
// @Description Create an attribute with the provided JSON payload
// @Tags Attribute
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
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.attributeService.CreateAttribute(&attribute, ctx); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, attribute)
}

// GetAttributeByID godoc
// @Summary Get Attribute by ID
// @Description Get a single attribute by its ID
// @Tags Attribute
// @Produce  json
// @Param   id  path  int  true  "Attribute ID"
// @Success 200 {object} models.Attribute
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/{id} [get]
func (c *AttributeHandler) GetAttributeByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid attribute ID"})
		return
	}

	attribute, err := c.attributeService.GetAttributeByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, attribute)
}

// GetAllAttributes godoc
// @Summary Get all attributes
// @Description Retrieve a list of all attributes.
// @Description - Use 'search' parameter for full-text search across all searchable fields.
// @Description - Use 'code', 'id', or 'name' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', and 'name' parameters for cross-field AND search.
// @Description - sort[any_field]=asc or sort[ant_field]=desc
// @Description Example: /?search=keyword&code=abc&name=test&sort[name]=asc&sort[id]=desc
// @Tags Attribute
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search term for full-text search across searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Produce json
// @Success 200 {array} models.Attribute
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/ [get]
func (c *AttributeHandler) GetAllAttributes(ctx *gin.Context) {
	attributes, err := c.attributeService.GetAllAttributes(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, attributes)
}

// UpdateAttribute godoc
// @Summary Update an existing attribute
// @Description Update the details of an existing attribute by providing the updated JSON payload
// @Tags Attribute
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Attribute ID"
// @Param   product  body  models.Attribute  true  "Updated Attribute data"
// @Success 200 {object} models.Attribute
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/{id} [put]
func (c *AttributeHandler) UpdateAttribute(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid attribute ID"})
		return
	}

	var attribute models.Attribute
	if err := ctx.ShouldBindJSON(&attribute); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	err = c.attributeService.UpdateAttribute(uint(id), &attribute, ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, attribute)
}

// DeleteAttribute godoc
// @Summary Delete an attribute by ID
// @Description Delete a single attribute by its ID
// @Tags Attribute
// @Produce json
// @Param id path int true "Attribute ID"
// @Success 200 {object} map[string]interface{} "Attribute deleted successfully"
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attributes/{id} [delete]
func (c *AttributeHandler) DeleteAttribute(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid attribute ID"})
		return
	}

	if err := c.attributeService.ArchiveAttribute(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Attribute deleted successfully", "id": id})
}
