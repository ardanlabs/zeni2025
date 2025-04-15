package web

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const (
	traceIDKey ctxKey = iota + 1
)

func setTraceID(ctx context.Context, traceID uuid.UUID) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

func GetTraceID(ctx context.Context) uuid.UUID {
	traceID, exists := ctx.Value(traceIDKey).(uuid.UUID)
	if !exists {
		return uuid.UUID{}
	}

	return traceID
}
