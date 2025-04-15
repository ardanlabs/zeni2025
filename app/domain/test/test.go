package test

import (
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

func (a app) handler(w http.ResponseWriter, r *http.Request) {
	a.log.Info(r.Context(), "handler", "path", r.URL.Path, "status", "started")
	defer a.log.Info(r.Context(), "handler", "path", r.URL.Path, "status", "completed")

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
