package test

import (
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func Routes(log *logger.Logger, app *web.App) {
	api := new(log)

	app.HandleFunc("GET /test", api.handler)
}
