package service

import (
	"context"

	l "github.com/gapidobri/otel-demo/internal/pkg/observability/logging"
	t "github.com/gapidobri/otel-demo/internal/pkg/observability/tracing"
)

type (
	Service struct{}
)

func NewService() Service {
	return Service{}
}

func (s Service) Get(ctx context.Context) string {
	ctx, span := t.Tracer.Start(ctx, "service2.Get")
	defer span.End()

	l.Logger.Ctx(ctx).Info("Get")

	return "Data from service 2"
}
