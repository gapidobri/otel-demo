package http

import (
	l "github.com/gapidobri/otel-demo/internal/pkg/observability/logging"
	"github.com/gapidobri/otel-demo/internal/pkg/observability/metrics"
	"github.com/gapidobri/otel-demo/internal/service-1/service"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
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
	if err := metrics.SetupGin(r); err != nil {
		l.Logger.Fatal("Failed to setup gin metrics", zap.Error(err))
	}

	registerRoutes(r, s.service)

	r.Run()

}
