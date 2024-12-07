package gclog

import (
	"context"
	"net/http"

	"github.com/najeira/gctrace"
	"go.uber.org/zap"
)

const (
	fieldsContextKey contextKey = iota
)

type (
	contextKey int
)

func ContextWithTrace(r *http.Request) context.Context {
	return ContextWithFields(r.Context(), gctrace.Fields(r)...)
}

func ContextWithFields(ctx context.Context, fields ...zap.Field) context.Context {
	if len(fields) <= 0 {
		return ctx
	}
	currentFields := fieldsFromContext(ctx)
	newFields := append(currentFields, fields...)
	return context.WithValue(ctx, fieldsContextKey, newFields)
}

func fieldsFromContext(ctx context.Context) []zap.Field {
	if v := ctx.Value(fieldsContextKey); v != nil {
		fs, _ := v.([]zap.Field)
		return fs
	}
	return nil
}
