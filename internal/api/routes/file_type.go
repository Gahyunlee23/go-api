package routes

import (
	handler "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func FileTypeRoutes(router *gin.Engine, handler *handler.FileTypeHandler) {
	fileTypeGroup := router.Group("/file-type")
	{
		fileTypeGroup.GET("/", handler.GetAllFileTypes)
		fileTypeGroup.POST("/", handler.CreateFileType)
		fileTypeGroup.PUT("/:id", handler.UpdateFileType)
		fileTypeGroup.GET(":id", handler.GetFileTypeByID)
		fileTypeGroup.DELETE("/:id", handler.DeleteFileType)
	}
}
