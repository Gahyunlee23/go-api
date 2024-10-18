package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	"main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DenyRuleHandler struct {
	denyRuleService services.DenyRuleService
}

func NewDenyRuleHandler(service services.DenyRuleService) *DenyRuleHandler {
	return &DenyRuleHandler{denyRuleService: service}
}

// CreateDenyRule godoc
// @Summary Create a new deny rule
// @Description Create a deny rule with the provided JSON payload
// @Tags DenyRule
// @Accept  json
// @Produce  json
// @Param   denyRule  body  models.DenyRule  true  "Deny Rule data"
// @Success 200 {object} models.DenyRule
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /deny-rules/ [post]
func (c *DenyRuleHandler) CreateDenyRule(ctx *gin.Context) {
	var denyRule models.DenyRule

	if err := ctx.ShouldBindJSON(&denyRule); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.denyRuleService.CreateDenyRule(&denyRule, ctx); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, denyRule)
}

// GetDenyRuleByID godoc
// @Summary Get Deny Rule by ID
// @Description Get a single Deny Rule by its ID
// @Tags DenyRule
// @Produce  json
// @Param   id  path  int  true  "Deny Rule ID"
// @Success 200 {object} models.DenyRule
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /deny-rules/{id} [get]
func (c *DenyRuleHandler) GetDenyRuleByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Deny Rule ID"})
		return
	}

	denyRule, err := c.denyRuleService.GetDenyRuleByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, denyRule)
}

// GetAllDenyRules godoc
// @Summary Get all deny rules
// @Description Retrieve a list of all deny rules.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use the 'code', 'id', or 'name' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', and 'name' parameters for cross-field AND search.
// @Description Example: /deny-rules?search=keyword&code=abc&name=test
// @Tags DenyRule
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Produce  json
// @Success 200 {array} models.DenyRule
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /deny-rules/ [get]
func (c *DenyRuleHandler) GetAllDenyRules(ctx *gin.Context) {
	denyRules, err := c.denyRuleService.GetAllDenyRules(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, denyRules)
}

// UpdateDenyRule godoc
// @Summary Update an existing deny rule
// @Description Update the details of an existing deny rule by providing the updated JSON payload
// @Tags DenyRule
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Deny Rule ID"
// @Param   denyRule  body  models.DenyRule  true  "Updated deny rule data"
// @Success 200 {object} models.DenyRule "Successfully updated the deny rule"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 400 {object} models.ErrorResponse "ID mismatch error"
// @Failure 400 {object} models.ErrorResponse "Invalid ID in request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /deny-rules/{id} [put]
func (c *DenyRuleHandler) UpdateDenyRule(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Deny Rule ID"})
		return
	}

	var denyRule models.DenyRule
	if err := ctx.ShouldBindJSON(&denyRule); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.denyRuleService.UpdateDenyRule(uint(id), &denyRule, ctx); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, denyRule)
}

// DeleteDenyRule godoc
// @Summary Delete a deny rule by ID
// @Description Delete a single deny rule by its ID
// @Tags DenyRule
// @Produce  json
// @Param   id  path  int  true  "Deny Rule ID"
// @Success 200 {object} map[string]interface{} "Deny Rule deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /deny-rules/{id} [delete]
func (c *DenyRuleHandler) DeleteDenyRule(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Deny Rule ID"})
		return
	}

	if err := c.denyRuleService.ArchiveDenyRule(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deny Rule deleted successfully", "id": id})
}
