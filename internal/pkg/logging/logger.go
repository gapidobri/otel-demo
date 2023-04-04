package logging

import (
	"os"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *otelzap.Logger {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		zapcore.DebugLevel,
	)
	return otelzap.New(zap.New(core))
}
