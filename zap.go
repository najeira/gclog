package google_cloud_log

import (
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

func zapNewProduction(serviceName string) (*zap.Logger, error) {
	cfg := zapdriver.NewDevelopmentConfig()
	cfg.Development = false
	return cfg.Build(zapCore(serviceName), zap.AddCallerSkip(1))
}

func zapNewDevelopment(serviceName string) (*zap.Logger, error) {
	return zapdriver.NewDevelopmentWithCore(zapCore(serviceName), zap.AddCallerSkip(1))
}

func zapCore(serviceName string) zap.Option {
	return zapdriver.WrapCore(
		zapdriver.ReportAllErrors(true),
		zapdriver.ServiceName(serviceName),
	)
}

func zapNewLocal(serviceName string) (*zap.Logger, error) {
	return zap.NewDevelopment(zap.Fields(
		zap.String("service", serviceName),
	), zap.AddCallerSkip(1))
}
