package test

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ardanlabs/service/foundation/logger"
)

type app struct {
	log *logger.Logger
}

func new(log *logger.Logger) *app {
	return &app{
		log: log,
	}
}

func (a app) handler(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Recv Input
	// Validate Input
	// Process OK response

	a.log.Info(ctx, "handler", "path", r.URL.Path, "status", "started")
	defer a.log.Info(ctx, "handler", "path", r.URL.Path, "status", "completed")

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return json.NewEncoder(w).Encode(status)
}
