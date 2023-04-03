package http

import (
	"github.com/gapidobri/otel-demo/internal/app/service"
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, s service.Service) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": s.Get(),
		})
	})
}
