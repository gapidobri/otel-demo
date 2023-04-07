package metrics

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument"
)

func SetupGin(r *gin.Engine) error {
	counter, err := Meter.Int64Counter(
		"http_requests_total",
		instrument.WithDescription("Total number of HTTP requests"),
	)
	if err != nil {
		return fmt.Errorf("failed to create counter: %w", err)
	}

	r.Use(func(c *gin.Context) {
		counter.Add(
			c, 1,
			attribute.Int("code", c.Writer.Status()),
			attribute.String("method", c.Request.Method),
			attribute.String("path", c.Request.URL.Path),
		)
		c.Next()
	})

	return nil
}
