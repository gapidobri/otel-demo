package service

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
)

type (
	Service struct {
		logger *otelzap.Logger
		tracer trace.Tracer
	}
)

func NewService(logger *otelzap.Logger, tracer trace.Tracer) Service {
	return Service{
		logger: logger,
		tracer: tracer,
	}
}

func (s Service) Get(ctx context.Context) string {
	ctx, span := s.tracer.Start(ctx, "service.Get")
	defer span.End()

	logger := s.logger.Ctx(ctx)

	logger.Info("Called get")

	return "Hello World!"
}
