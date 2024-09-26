package controllers

import (
	"fmt"
	"main-admin-api/models"
	"main-admin-api/services"
	"main-admin-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DenyRuleController struct {
	denyRuleService services.DenyRuleServiceInterface
}

func NewDenyRuleController(service services.DenyRuleServiceInterface) *DenyRuleController {
	return &DenyRuleController{denyRuleService: service}
}

// CreateDenyRule godoc
// @Summary Create a new deny rule
// @Description Create a deny rule with the provided JSON payload
// @Tags DenyRule
// @Accept  json
// @Produce  json
// @Param   denyRule  body  models.DenyRule  true  "Deny Rule data"
// @Success 200 {object} models.DenyRule
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /deny-rules/ [post]
func (c *DenyRuleController) CreateDenyRule(ctx *gin.Context) {
	var denyRule models.DenyRule

	if err := ctx.ShouldBindJSON(&denyRule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// when the value is json type, should bind make them to Go struct but convert json again to insert
	var err error
	denyRule.Paper, err = utils.MarshalAndAssignJSON(denyRule.Paper, "paper", ctx)
	if err != nil {
		return
	}

	denyRule.Format, err = utils.MarshalAndAssignJSON(denyRule.Format, "format", ctx)
	if err != nil {
		return
	}

	denyRule.Pages, err = utils.MarshalAndAssignJSON(denyRule.Pages, "pages", ctx)
	if err != nil {
		return
	}

	denyRule.Colors, err = utils.MarshalAndAssignJSON(denyRule.Colors, "colors", ctx)
	if err != nil {
		return
	}

	denyRule.BookBinding, err = utils.MarshalAndAssignJSON(denyRule.BookBinding, "bookBinding", ctx)
	if err != nil {
		return
	}

	denyRule.Refinement, err = utils.MarshalAndAssignJSON(denyRule.Refinement, "refinement", ctx)
	if err != nil {
		return
	}

	denyRule.Finishing, err = utils.MarshalAndAssignJSON(denyRule.Finishing, "finishing", ctx)
	if err != nil {
		return
	}

	if err := c.denyRuleService.CreateDenyRule(&denyRule); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"denyRule": denyRule})

}

// GetDenyRuleByID godoc
// @Summary Get Deny Rule by ID
// @Description Get a single Deny Rule by its ID
// @Tags DenyRule
// @Produce  json
// @Param   id  path  int  true  "Deny Rule ID"
// @Success 200 {object} models.DenyRule
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Deny Rule not found"
// @Router /deny-rules/{id} [get]
func (c *DenyRuleController) GetDenyRuleByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid ID": err.Error()})
	}
	denyRule, err := c.denyRuleService.GetDenyRuleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"denyRule": denyRule})
}

// GetAllDenyRules GetDenyRules godoc
// @Summary Get all deny rules
// @Description Retrieve a list of all deny rules
// @Tags DenyRule
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search term for filtering by name or code"
// @Produce  json
// @Success 200 {array} models.DenyRule
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /deny-rules/ [get]
func (c *DenyRuleController) GetAllDenyRules(ctx *gin.Context) {
	denyRules, err := c.denyRuleService.GetAllDenyRules(ctx)
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	fmt.Println("Page:", page, "Page Size:", pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"denyRules": denyRules})
}

// UpdateDenyRule godoc
// @Summary Update an existing deny rule
// @Description Update the details of an existing deny rule by providing the updated JSON payload
// @Tags DenyRule
// @Accept  json
// @Produce  json
// @Param   denyRule  body  models.DenyRule  true  "Updated deny rule data"
// @Success 200 {object} models.DenyRule
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /deny-rules/{id} [put]
func (c *DenyRuleController) UpdateDenyRule(ctx *gin.Context) {
	var denyRule models.DenyRule
	if err := ctx.ShouldBindJSON(&denyRule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	denyRule.Paper, err = utils.MarshalAndAssignJSON(denyRule.Paper, "paper", ctx)
	if err != nil {
		return
	}

	denyRule.Format, err = utils.MarshalAndAssignJSON(denyRule.Format, "format", ctx)
	if err != nil {
		return
	}

	denyRule.Pages, err = utils.MarshalAndAssignJSON(denyRule.Pages, "pages", ctx)
	if err != nil {
		return
	}

	denyRule.Colors, err = utils.MarshalAndAssignJSON(denyRule.Colors, "colors", ctx)
	if err != nil {
		return
	}

	denyRule.BookBinding, err = utils.MarshalAndAssignJSON(denyRule.BookBinding, "bookBinding", ctx)
	if err != nil {
		return
	}

	denyRule.Refinement, err = utils.MarshalAndAssignJSON(denyRule.Refinement, "refinement", ctx)
	if err != nil {
		return
	}

	denyRule.Finishing, err = utils.MarshalAndAssignJSON(denyRule.Finishing, "finishing", ctx)
	if err != nil {
		return
	}

	if err := c.denyRuleService.UpdateDenyRule(&denyRule); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"denyRule": denyRule})
}

// DeleteDenyRule godoc
// @Summary Delete a deny rule by ID
// @Description Delete a single deny rule by its ID
// @Tags DenyRule
// @Produce  json
// @Param   id  path  int  true  "Deny Rule ID"
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /deny-rules/{id} [delete]
func (c *DenyRuleController) DeleteDenyRule(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Invalid ID": err.Error()})
	}
	if err := c.denyRuleService.DeleteDenyRule(uint(id)); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"denyRule": nil})
}
