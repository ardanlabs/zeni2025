// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"net/http"

	"github.com/ardanlabs/service/app/domain/test"
	"github.com/ardanlabs/service/foundation/logger"
)

func WebAPI(log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	test.Routes(log, mux)

	return mux
}
