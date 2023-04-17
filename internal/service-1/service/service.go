package service

import (
	"context"

	l "github.com/gapidobri/otel-demo/internal/pkg/observability/logging"
	t "github.com/gapidobri/otel-demo/internal/pkg/observability/tracing"
	service2Client "github.com/gapidobri/otel-demo/internal/service-2/pkg/rabbit"
)

type (
	Service struct {
		service2Client service2Client.Client
	}
)

func NewService(service2Client service2Client.Client) Service {
	return Service{
		service2Client: service2Client,
	}
}

func (s Service) Get(ctx context.Context) string {
	ctx, span := t.Tracer.Start(ctx, "service1.Get")
	defer span.End()

	logger := l.Logger.Ctx(ctx)

	logger.Info("Called get")

	return "Hello World!"
}

func (s Service) GetService2(ctx context.Context) string {
	ctx, span := t.Tracer.Start(ctx, "service1.GetService2")
	defer span.End()

	logger := l.Logger.Ctx(ctx)

	err := s.service2Client.Get(ctx)
	if err != nil {
		return err.Error()
	}

	logger.Info("Called get service 2")

	return "Hello World!"
}
