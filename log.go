package gclog

import (
	"context"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	logger = zap.NewNop()
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, fieldsFromContext(ctx)...)
	logger.Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, fieldsFromContext(ctx)...)
	logger.Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, fieldsFromContext(ctx)...)
	logger.Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, fieldsFromContext(ctx)...)
	logger.Error(msg, fields...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, fieldsFromContext(ctx)...)
	logger.Panic(msg, fields...)
}

func SetProduction(serviceName string) error {
	return setLogger(zapNewProduction(serviceName))
}

func SetDevelopment(serviceName string) error {
	return setLogger(zapNewDevelopment(serviceName))
}

func SetLocal(serviceName string) error {
	return setLogger(zapNewLocal(serviceName))
}

func setLogger(l *zap.Logger, err error) error {
	if err != nil {
		return err
	}
	logger = l
	return nil
}
