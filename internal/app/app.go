package app

import (
	"context"

	"github.com/gapidobri/otel-demo/internal/app/api/http"
	"github.com/gapidobri/otel-demo/internal/app/service"
	"github.com/gapidobri/otel-demo/internal/pkg/logging"
	"github.com/gapidobri/otel-demo/internal/pkg/metrics"
	"github.com/gapidobri/otel-demo/internal/pkg/tracing"
	"go.uber.org/zap"
)

func Run() {
	ctx := context.Background()

	// Tracing
	{
		shutdown, err := tracing.Setup(ctx)
		if err != nil {
			logging.Logger.Fatal("Failed to setup tracing", zap.Error(err))
		}
		defer shutdown(ctx)
	}

	// Metrics
	{
		shutdown, err := metrics.Setup(ctx)
		if err != nil {
			logging.Logger.Fatal("Failed to setup metrics", zap.Error(err))
		}
		defer shutdown(ctx)
	}

	service := service.NewService()

	server := http.NewServer(service)
	server.Run()
}
