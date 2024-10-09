package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProofHandler struct {
	proofService services.ProofService
}

func NewProofHandler(service services.ProofService) *ProofHandler {
	return &ProofHandler{proofService: service}
}

// CreateProof godoc
// @Summary Create a new Proof
// @Description Create an Proof with the provided JSON payload
// @Tags Proofs
// @Accept  json
// @Produce  json
// @Param   Proof  body  models.Proof  true  "Proof data"
// @Success 200 {object} models.Proof
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /proofs/ [post]
func (c *ProofHandler) CreateProof(ctx *gin.Context) {
	var Proof models.Proof

	if err := ctx.ShouldBindJSON(&Proof); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.proofService.CreateProof(&Proof); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, Proof)
}

// GetProofByID godoc
// @Summary Get Proof by ID
// @Description Get a single Proof by its ID
// @Tags Proofs
// @Produce  json
// @Param   id  path  int  true  "Proof ID"
// @Success 200 {object} models.Proof
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /proofs/{id} [get]
func (c *ProofHandler) GetProofByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Proof ID"})
		return
	}

	proof, err := c.proofService.GetProofByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, proof)
}

// GetAllProofs godoc
// @Summary Get all Proofs
// @Description Retrieve a list of all Proofs.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use the 'code', 'id', 'name', or 'type' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', 'name', and 'type' parameters for cross-field AND search.
// @Description Example: /proofs?search=keyword&code=abc&name=test
// @Tags Proofs
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Param type query string false "Filter by type field (partial match)"
// @Produce  json
// @Success 200 {array} models.Proof
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /proofs/ [get]
func (c *ProofHandler) GetAllProofs(ctx *gin.Context) {
	proofs, err := c.proofService.GetAllProofs(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, proofs)
}

// UpdateProof godoc
// @Summary Update an existing Proof
// @Description Update the details of an existing Proof by providing the updated JSON payload
// @Tags Proofs
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "Proof ID"
// @Param   Proof  body  models.Proof  true  "Updated Proof data"
// @Success 200 {object} models.Proof
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /Proofs/{id} [put]
func (c *ProofHandler) UpdateProof(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Proof ID"})
		return
	}

	var proof models.Proof
	if err := ctx.ShouldBindJSON(&proof); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.proofService.UpdateProof(uint(id), &proof); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, proof)
}

// DeleteProof godoc
// @Summary Delete a Proof by ID
// @Description Delete a single Proof by its ID
// @Tags Proofs
// @Produce  json
// @Param   id  path  int  true  "Proof ID"
// @Success 200 {object} map[string]interface{} "Proof deleted successfully"
// @Failure 400 {object} map[string]interface{} "Validation error on field '%Given ID'"
// @Failure 404 {object} map[string]interface{} "Entity '%Entity Type' with ID '%Given ID' not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /Proofs/{id} [delete]
func (c *ProofHandler) DeleteProof(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid Proof ID"})
		return
	}

	if err := c.proofService.ArchiveProof(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Proof deleted successfully", "id": id})
}
