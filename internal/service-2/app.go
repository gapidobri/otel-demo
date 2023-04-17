package service_2

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/GLCharge/backend-services/pkg/models/configuration"
	"github.com/GLCharge/backend-services/pkg/rabbit"
	l "github.com/gapidobri/otel-demo/internal/pkg/observability/logging"
	rabbitApi "github.com/gapidobri/otel-demo/internal/service-2/api/rabbit"
	"github.com/gapidobri/otel-demo/internal/service-2/service"
	"go.uber.org/zap"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := l.Logger.Ctx(ctx)

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

	// Service
	service := service.NewService()

	rabbitApi := rabbitApi.NewAPI(service, rabbit)
	rabbitApi.Start()

	<-ctx.Done()
	logger.Info("Stopping service 2")
}
