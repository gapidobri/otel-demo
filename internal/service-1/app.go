package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/GLCharge/backend-services/pkg/models/configuration"
	"github.com/GLCharge/backend-services/pkg/rabbit"
	"github.com/gapidobri/otel-demo/internal/pkg/observability/logging"
	l "github.com/gapidobri/otel-demo/internal/pkg/observability/logging"
	"github.com/gapidobri/otel-demo/internal/pkg/observability/metrics"
	"github.com/gapidobri/otel-demo/internal/pkg/observability/tracing"
	"github.com/gapidobri/otel-demo/internal/service-1/api/http"
	"github.com/gapidobri/otel-demo/internal/service-1/service"
	service2Client "github.com/gapidobri/otel-demo/internal/service-2/pkg/rabbit"
	"go.uber.org/zap"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := l.Logger.Ctx(ctx)

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

	// RabbitMQ
	rabbitConfig := rabbit.GetConfiguration(
		configuration.Rabbit{
			Address:  "localhost",
			Username: "guest",
			Password: "guest",
		},
		rabbit.Exchange("service_2"),
	)

	rabbit, err := rabbit.NewRabbit(
		rabbitConfig,
		rabbit.WithOptionsConcurrentReplyConsumer(2),
		rabbit.WithOptionsMultiplePublishers(2),
	)
	if err != nil {
		logger.Error("Could not create rabbit client", zap.Error(err))
	}
	defer rabbit.Disconnect()

	service2Client := service2Client.NewClient(rabbit)

	// Service
	service := service.NewService(service2Client)

	server := http.NewServer(service)
	server.Run()
}
