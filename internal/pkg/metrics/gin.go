package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func SetupGin(r *gin.Engine) {
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
	}, []string{"code", "method", "path"})

	AddCollector(counter)

	r.Use(func(c *gin.Context) {
		counter.WithLabelValues(
			strconv.Itoa(c.Writer.Status()),
			c.Request.Method,
			c.Request.URL.Path,
		).Inc()

		c.Next()
	})

	go func() {
		for {
			pusher.Push()
			time.Sleep(5 * time.Second)
		}
	}()
}
