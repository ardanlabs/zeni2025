package test

import (
	"net/http"

	"github.com/ardanlabs/service/foundation/logger"
)

func Routes(log *logger.Logger, mux *http.ServeMux) {
	api := new(log)

	mux.HandleFunc("GET /test", api.handler)
}
