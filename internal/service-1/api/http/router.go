package http

import (
	"github.com/gapidobri/otel-demo/internal/service-1/service"
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

	r.GET("/service-2", func(c *gin.Context) {
		ctx := c.Request.Context()
		data := s.GetService2(ctx)
		c.JSON(200, gin.H{
			"message": data,
		})
	})
}
