package rabbit

import (
	"context"
	"time"

	"github.com/GLCharge/backend-services/pkg/rabbit"
	"github.com/gapidobri/otel-demo/internal/service-2/service"
	"github.com/wagslane/go-rabbitmq"
)

type Api struct {
	rabbit  *rabbit.Rabbit
	service service.Service
}

func NewAPI(service service.Service, rb *rabbit.Rabbit) rabbit.Api {
	rabbitApi := &Api{
		rabbit:  rb,
		service: service,
	}
	return rabbitApi
}

func (a *Api) Start() {
	a.GetConsumer()
}

func (a *Api) GetConsumer() {
	handlerFunc := func(d rabbitmq.Delivery) (action rabbitmq.Action) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		res := a.service.Get(ctx)

		a.rabbit.Publisher.Respond(ctx, d.CorrelationId, rabbit.Topic(d.ReplyTo), []byte(res), false)

		return rabbitmq.Ack
	}

	topic := rabbit.NewTopic(a.rabbit.Exchange).AddWord("get").Build()
	a.rabbit.Consumer.NewTopicConsumer(topic, handlerFunc, true, 1)
}
