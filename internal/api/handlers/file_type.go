package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileTypeHandler struct {
	fileTypeService services.FileTypeService
}

func NewFileTypeHandler(service services.FileTypeService) *FileTypeHandler {
	return &FileTypeHandler{fileTypeService: service}
}

// CreateFileType godoc
// @Summary Create a new File Type
// @Description Create a File Type with the provided JSON payload
// @Tags FileType
// @Accept  json
// @Produce  json
// @Param   FileType  body  models.FileType  true  "File Type data"
// @Success 200 {object} models.FileType
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-types/ [post]
func (c *FileTypeHandler) CreateFileType(ctx *gin.Context) {
	var fileType models.FileType

	if err := ctx.ShouldBindJSON(&fileType); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.fileTypeService.CreateFileType(&fileType); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, fileType)
}

// GetFileTypeByID godoc
// @Summary Get File Type by ID
// @Description Get a single File Type by its ID
// @Tags FileType
// @Produce  json
// @Param   id  path  int  true  "File Type ID"
// @Success 200 {object} models.FileType
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-types/{id} [get]
func (c *FileTypeHandler) GetFileTypeByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid File Type ID"})
		return
	}

	fileType, err := c.fileTypeService.GetFileTypeByID(uint(id))
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fileType)
}

// GetAllFileTypes godoc
// @Summary Get all file types
// @Description Retrieve a list of all file types.
// @Description - Use the 'search' parameter for a full-text search across all searchable fields.
// @Description - Use the 'code', 'id', or 'name' parameters for individual field searches (partial matches).
// @Description - Combine 'code', 'id', and 'name' parameters for cross-field AND search.
// @Description Example: /file types?search=keyword&code=abc&name=test
// @Tags FileType
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Full-text search across all searchable fields"
// @Param code query string false "Filter by code field (partial match)"
// @Param id query string false "Filter by ID field (partial match)"
// @Param name query string false "Filter by name field (partial match)"
// @Produce  json
// @Success 200 {array} models.FileType
// / @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-types/ [get]
func (c *FileTypeHandler) GetAllFileTypes(ctx *gin.Context) {
	fileType, err := c.fileTypeService.GetAllFileTypes(ctx)
	if err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fileType)
}

// UpdateFileType godoc
// @Summary Update an existing FileType
// @Description Update the details of an existing File Type by providing the updated JSON payload
// @Tags FileType
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "FileType ID"
// @Param   product  body  models.FixedPrice  true  "Updated File Type data"
// @Success 200 {object} models.FixedPrice
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 400 {object} models.ErrorResponse "ID mismatch error"
// @Failure 400 {object} models.ErrorResponse "Invalid ID in request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /file-types/{id} [put]
func (c *FileTypeHandler) UpdateFileType(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid File Type ID"})
		return
	}

	var fileType models.FileType
	if err := ctx.ShouldBindJSON(&fileType); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.fileTypeService.UpdateFileType(uint(id), &fileType); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, fileType)
}

// DeleteFileType godoc
// @Summary Delete a File Type by ID
// @Description Delete a single File Type by its ID
// @Tags FileType
// @Produce json
// @Param id path int true "File Type ID"
// @Success 200 {object} map[string]interface{} "File Type deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 404 {object} models.ErrorResponse "Entity not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /file-types/{id} [delete]
func (c *FileTypeHandler) DeleteFileType(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "id", Message: "Invalid File Type ID"})
		return
	}

	if err := c.fileTypeService.ArchiveFileType(uint(id)); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File Type deleted successfully", "id": id})
}
