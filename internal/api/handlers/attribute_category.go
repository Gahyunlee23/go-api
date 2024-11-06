package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttributeCategoryHandler struct {
	attributeCategoryService services.AttributeCategoryService
}

func NewAttributeCategory(service services.AttributeCategoryService) *AttributeCategoryHandler {
	return &AttributeCategoryHandler{attributeCategoryService: service}
}

// GetAttributeCategoryByID godoc
// @Summary Get AttributeCategory by ID
// @Description Get a single AttributeCategory by its ID
// @Tags AttributeCategory
// @Produce  json
// @Param   id  path  int  true  "AttributeCategory ID"
// @Success 200 {object} models.AttributeCategory
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attribute-categories/{id} [get]
func (c *AttributeCategoryHandler) GetAttributeCategoryByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Attribute Category ID"})
		return
	}

	AttributeCategory, err := c.attributeCategoryService.GetAttributesCategoryByID(uint(id), ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, AttributeCategory)
}

// GetAllAttributeCategory godoc
// @Summary Get all AttributeCategory
// @Description Retrieve a list of all AttributeCategory.
// @Description - Use 'search' parameter for full-text search across all searchable fields.
// @Description - Use 'code', 'id', or 'name' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', and 'name' parameters for cross-field AND search.
// @Description - sort[any_field]=asc or sort[ant_field]=desc
// @Description Example: /?search=keyword&code=abc&name=test&sort[name]=asc&sort[id]=desc
// @Tags AttributeCategory
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search term for full-text search across searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Produce json
// @Success 200 {array} models.AttributeCategory
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attribute-categories/ [get]
func (c *AttributeCategoryHandler) GetAllAttributeCategory(ctx *gin.Context) {
	attributeCategory, err := c.attributeCategoryService.GetAllAttributesCategories(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, attributeCategory)
}

// CreateAttributeCategory godoc
// @Summary Create a new AttributeCategory
// @Description Create an AttributeCategory with the provided JSON payload
// @Tags AttributeCategory
// @Accept  json
// @Produce  json
// @Param   AttributeCategory  body  models.AttributeCategory  true  "Attribute Category data"
// @Success 200 {object} models.AttributeCategory
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attribute-categories/ [post]
func (c *AttributeCategoryHandler) CreateAttributeCategory(ctx *gin.Context) {
	var attributeCategory models.AttributeCategory

	if err := ctx.ShouldBindJSON(&attributeCategory); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.attributeCategoryService.CreateAttributeCategory(&attributeCategory); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, attributeCategory)
}

// UpdateAttributeCategory godoc
// @Summary Update an existing AttributeCategory
// @Description Update the details of an existing AttributeCategory by providing the updated JSON payload
// @Tags AttributeCategory
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Attribute Category ID"
// @Param   product  body  models.AttributeCategory  true  "Updated Attribute Category data"
// @Success 200 {object} models.AttributeCategory
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attribute-categories/{id} [put]
func (c *AttributeCategoryHandler) UpdateAttributeCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid AttributeCategory ID"})
		return
	}

	var attributeCategory models.AttributeCategory
	if err := ctx.ShouldBindJSON(&attributeCategory); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	err = c.attributeCategoryService.UpdateAttributeCategory(uint(id), &attributeCategory)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, attributeCategory)
}

// DeleteAttributeCategory godoc
// @Summary Delete an Attribute Category by ID
// @Description Delete a single Attribute Category by its ID
// @Tags AttributeCategory
// @Produce json
// @Param id path int true "Attribute Category ID"
// @Success 200 {object} map[string]interface{} "Attribute Category deleted successfully"
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /attribute-categories/{id} [delete]
func (c *AttributeCategoryHandler) DeleteAttributeCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Attribute Category ID"})
		return
	}

	if err := c.attributeCategoryService.ArchiveAttributeCategory(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Attribute Category deleted successfully", "id": id})
}
