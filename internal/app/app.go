package app

import (
	"context"

	"github.com/gapidobri/otel-demo/internal/app/api/http"
	"github.com/gapidobri/otel-demo/internal/app/service"
	"github.com/gapidobri/otel-demo/internal/pkg/logging"
	"github.com/gapidobri/otel-demo/internal/pkg/metrics"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"go.opentelemetry.io/otel"
)

func Run() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	// Logging
	logger := logging.NewLogger()

	// Tracing
	tracer := otel.Tracer("demo-service")

	// Metrics
	metrics.SetupMetrics()

	service := service.NewService(logger, tracer)

	server := http.NewServer(service)
	server.Run()
}
