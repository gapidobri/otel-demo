package http

import (
	"github.com/gapidobri/otel-demo/internal/app/service"
	"github.com/gapidobri/otel-demo/internal/pkg/metrics"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type (
	Server struct {
		service service.Service
	}
)

func NewServer(service service.Service) Server {
	return Server{
		service: service,
	}
}

func (s Server) Run() {
	r := gin.Default()

	// Tracing
	r.Use(otelgin.Middleware("demo-service"))

	// Metrics
	metrics.SetupGin(r)

	registerRoutes(r, s.service)

	r.Run()

}
