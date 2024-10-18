package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SelectionRuleHandler struct {
	selectionRuleService services.SelectionRuleService
}

func NewSelectionRuleHandler(service services.SelectionRuleService) *SelectionRuleHandler {
	return &SelectionRuleHandler{selectionRuleService: service}
}

// CreateSelectionRule godoc
// @Summary Create a new selection rule
// @Description Create a selection rule with the provided JSON payload
// @Tags SelectionRule
// @Accept  json
// @Produce  json
// @Param   SelectionRule  body  models.SelectionRule  true  "Selection Rule data"
// @Success 200 {object} models.SelectionRule
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /selection-rules/ [post]
func (c *SelectionRuleHandler) CreateSelectionRule(ctx *gin.Context) {
	var selectionRule models.SelectionRule

	if err := ctx.ShouldBindJSON(&selectionRule); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.selectionRuleService.CreateSelectionRule(ctx, &selectionRule); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, selectionRule)
}

// GetSelectionRuleByID godoc
// @Summary Get SelectionRule by ID
// @Description Get a single Selection Rule by its ID
// @Tags SelectionRule
// @Produce  json
// @Param   id  path  int  true  "Selection Rule ID"
// @Success 200 {object} models.SelectionRule
// @Failure 400 {object} map[string]interface{} "Entity 'Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /selection-rules/{id} [get]
func (c *SelectionRuleHandler) GetSelectionRuleByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: err.Error()})
		return
	}

	selectionRule, err := c.selectionRuleService.GetSelectionRuleByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, selectionRule)
}

// GetAllSelectionRules godoc
// @Summary Get all selection rules
// @Description Retrieve a list of all selection rules.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use 'code' or 'id', or 'name' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', and 'name' parameters for cross-field AND search.
// @Description Example: /selection-rules?search=keyword&code=abc&name=test
// @Tags SelectionRule
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Produce  json
// @Success 200 {array} models.SelectionRule
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /selection-rules/ [get]
func (c *SelectionRuleHandler) GetAllSelectionRules(ctx *gin.Context) {
	selectionRules, err := c.selectionRuleService.GetAllSelectionRules(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, selectionRules)
}

// UpdateSelectionRule godoc
// @Summary Update an existing selection rule
// @Description Update the details of an existing selection rule by providing the updated JSON payload
// @Tags SelectionRule
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Selection Rule ID"
// @Param   SelectionRule  body  models.SelectionRule  true  "Updated Selection Rule data"
// @Success 200 {object} models.SelectionRule
// @Failure 400 {object} map[string]interface{} "Entity 'Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /selection-rules/{id} [put]
func (c *SelectionRuleHandler) UpdateSelectionRule(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid SelectionRule ID"})
		return
	}

	var SelectionRule models.SelectionRule
	if err := ctx.ShouldBindJSON(&SelectionRule); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.selectionRuleService.UpdateSelectionRule(uint(id), &SelectionRule, ctx); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, SelectionRule)
}

// DeleteSelectionRule godoc
// @Summary Delete a selection rule by ID
// @Description Delete a single selection rule by its ID
// @Tags SelectionRule
// @Produce  json
// @Param   id  path  int  true  "Selection Rule ID"
// @Success 200 {object} map[string]interface{} "SelectionRule deleted successfully"
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /selection-rules/{id} [delete]
func (c *SelectionRuleHandler) DeleteSelectionRule(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid SelectionRule ID"})
		return
	}

	if err := c.selectionRuleService.ArchiveSelectionRule(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Selection Rule deleted successfully", "id": id})
}
