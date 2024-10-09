package routes

import (
	handler "main-admin-api/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func ProofRoute(router *gin.Engine, handler *handler.ProofHandler) {
	proofGroup := router.Group("/proof")
	{
		proofGroup.GET("/", handler.GetAllProofs)
		proofGroup.POST("/", handler.CreateProof)
		proofGroup.GET("/:id", handler.GetProofByID)
		proofGroup.PUT("/:id", handler.UpdateProof)
		proofGroup.DELETE("/:id", handler.DeleteProof)
	}
}
