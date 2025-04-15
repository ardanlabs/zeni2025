package test

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

type app struct {
	log *logger.Logger
}

func new(log *logger.Logger) *app {
	return &app{
		log: log,
	}
}

func (a app) handler(ctx context.Context, r *http.Request) web.Encoder {
	// Recv Input
	// Validate Input
	// Process OK response

	status := status{
		Status: "OK",
	}

	return status
}
