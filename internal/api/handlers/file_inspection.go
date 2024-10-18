package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileInspectionHandler struct {
	fileInspectionService services.FileInspectionService
}

func NewFileInspectionHandler(service services.FileInspectionService) *FileInspectionHandler {
	return &FileInspectionHandler{fileInspectionService: service}
}

// CreateFileInspection godoc
// @Summary Create a new File Inspection
// @Description Create a File Inspection with the provided JSON payload
// @Tags FileInspection
// @Accept  json
// @Produce  json
// @Param   FileInspection  body  models.FileInspection  true  "File Inspection data"
// @Success 200 {object} models.FileInspection
// / @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-inspections/ [post]
func (c *FileInspectionHandler) CreateFileInspection(ctx *gin.Context) {
	var fileInspection models.FileInspection

	if err := ctx.ShouldBindJSON(&fileInspection); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.fileInspectionService.CreateFileInspection(&fileInspection); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, fileInspection)
}

// GetFileInspectionByID godoc
// @Summary Get File Inspection by ID
// @Description Get a single File Inspection by its ID
// @Tags FileInspection
// @Produce  json
// @Param   id  path  int  true  "File Inspection ID"
// @Success 200 {object} models.FileInspection
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-inspections/{id} [get]
func (c *FileInspectionHandler) GetFileInspectionByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid File Inspection ID"})
		return
	}

	fileInspection, err := c.fileInspectionService.GetFileInspectionByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fileInspection)
}

// GetAllFileInspections godoc
// @Summary Get all file Inspections
// @Description Retrieve a list of all file Inspections.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use the 'code', 'id', or 'name' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', and 'name' parameters for cross-field AND search.
// @Description Example: /file Inspections?search=keyword&code=abc&name=test
// @Tags FileInspection
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Produce  json
// @Success 200 {array} models.FileInspection
// / @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-inspections/ [get]
func (c *FileInspectionHandler) GetAllFileInspections(ctx *gin.Context) {
	fileInspection, err := c.fileInspectionService.GetAllFileInspections(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fileInspection)
}

// UpdateFileInspection godoc
// @Summary Update an existing FileInspection
// @Description Update the details of an existing File Inspection by providing the updated JSON payload
// @Tags FileInspection
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "FileInspection ID"
// @Param   product  body  models.FixedPrice  true  "Updated File Inspection data"
// @Success 200 {object} models.FixedPrice "successfully updated fixed price"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 400 {object} models.ErrorResponse "ID mismatch error"
// @Failure 400 {object} models.ErrorResponse "Invalid ID in request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-inspections/{id} [put]
func (c *FileInspectionHandler) UpdateFileInspection(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid File Inspection ID"})
		return
	}

	var fileInspection models.FileInspection
	if err := ctx.ShouldBindJSON(&fileInspection); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.fileInspectionService.UpdateFileInspection(uint(id), &fileInspection); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fileInspection)
}

// DeleteFileInspection godoc
// @Summary Delete a File Inspection by ID
// @Description Delete a single File Inspection by its ID
// @Tags FileInspection
// @Produce json
// @Param id path int true "File Inspection ID"
// @Success 200 {object} map[string]interface{} "File Inspection deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-inspections/{id} [delete]
func (c *FileInspectionHandler) DeleteFileInspection(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid File Inspection ID"})
		return
	}

	if err := c.fileInspectionService.ArchiveFileInspection(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File Inspection deleted successfully", "id": id})
}
