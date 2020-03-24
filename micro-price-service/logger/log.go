package logger

import (
	"context"

	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ServiceName = "micro-price"
)

var (
	zLogger *zap.Logger
)

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.StacktraceKey = "error.stack_trace"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	config.OutputPaths = []string{
		"/var/logs/micro-price.log",
	}

	zLogger, _ = config.Build(zap.WrapCore((&apmzap.Core{}).WrapCore),
		zap.Fields(zap.String("service", ServiceName)))
}

//S Return zap sugared logger
func S() *zap.SugaredLogger {
	return zLogger.Sugar()
}

//Ctx Return context-aware logger for log correlation
func Ctx(ctx context.Context) *zap.Logger {
	traceContext := apmzap.TraceContext(ctx)
	if traceContext != nil {
		return zLogger.With(traceContext...)
	}
	return zLogger
}
