package middleware

import (
	"main-admin-api/pkg/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(cfg *config.Config) gin.HandlerFunc {
	maxAge, err := time.ParseDuration(cfg.CORS.MaxAge)
	if err != nil {
		maxAge = 12 * time.Hour // 기본값 설정
	}

	corsConfig := cors.Config{
		AllowMethods:     cfg.CORS.AllowMethods,
		AllowHeaders:     cfg.CORS.AllowHeaders,
		ExposeHeaders:    cfg.CORS.ExposeHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
		MaxAge:           maxAge,
	}

	if len(cfg.CORS.AllowOrigins) == 1 && cfg.CORS.AllowOrigins[0] == "*" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = cfg.CORS.AllowOrigins
	}

	return cors.New(corsConfig)
}
