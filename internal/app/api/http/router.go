package http

import (
	"github.com/gapidobri/otel-demo/internal/app/service"
	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, s service.Service) {
	r.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		data := s.Get(ctx)
		c.JSON(200, gin.H{
			"message": data,
		})
	})
}
