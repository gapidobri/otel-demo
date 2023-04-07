package service

import (
	"context"

	l "github.com/gapidobri/otel-demo/internal/pkg/logging"
	t "github.com/gapidobri/otel-demo/internal/pkg/tracing"
)

type (
	Service struct {
	}
)

func NewService() Service {
	return Service{}
}

func (s Service) Get(ctx context.Context) string {
	ctx, span := t.Tracer.Start(ctx, "service.Get")
	defer span.End()

	logger := l.Logger.Ctx(ctx)

	logger.Info("Called get")

	return "Hello World!"
}
