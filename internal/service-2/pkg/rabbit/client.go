package rabbit

import (
	"context"

	"github.com/GLCharge/backend-services/pkg/rabbit"
	"github.com/wagslane/go-rabbitmq"
	"go.opentelemetry.io/otel/trace"
)

const ServiceName = "service_2"

type (
	Client interface {
		Get(ctx context.Context) error
	}
	ClientImpl struct {
		rabbit *rabbit.Rabbit
	}
)

func NewClient(rabbit *rabbit.Rabbit) Client {
	return &ClientImpl{
		rabbit: rabbit,
	}
}

func (c *ClientImpl) RegisterConsumer(handler rabbitmq.Handler) (*rabbitmq.Consumer, error) {
	topic := rabbit.NewTopic(ServiceName).Build()
	return c.rabbit.Consumer.NewTopicConsumer(topic, handler, false, 1)
}

func (c *ClientImpl) Get(ctx context.Context) error {
	spanCtx := trace.SpanContextFromContext(ctx)
	headers := []rabbit.HeaderValue{
		{Key: "traceId", Value: spanCtx.TraceID().String()},
		{Key: "spanId", Value: spanCtx.SpanID().String()},
	}

	topic := rabbit.NewTopic(ServiceName).AddWord("get").Build()
	return c.rabbit.Publisher.Publish(ctx, topic, nil, headers)
}
