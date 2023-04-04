package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	pusher *push.Pusher
)

func SetupMetrics() {
	pusher = push.New("http://localhost:9091", "demo-service").
		Collector(collectors.NewGoCollector())

	go func() {
		for {
			pusher.Push()
			time.Sleep(5 * time.Second)
		}
	}()
}

func AddCollector(c prometheus.Collector) {
	pusher = pusher.Collector(c)
}
