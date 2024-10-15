package routes

import (
	Handlers "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func FileInspectionRoutes(router *gin.Engine, handler *Handlers.FileInspectionHandler) {
	FileInspectionGroup := router.Group("/file-inspections")
	{
		FileInspectionGroup.GET("/", handler.GetAllFileInspections)
		FileInspectionGroup.POST("/", handler.GetAllFileInspections)
		FileInspectionGroup.GET(":id", handler.GetFileInspectionByID)
		FileInspectionGroup.PUT("/:id", handler.UpdateFileInspection)
		FileInspectionGroup.DELETE("/:id", handler.DeleteFileInspection)
	}
}
