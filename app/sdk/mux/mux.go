// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"github.com/ardanlabs/service/app/domain/test"
	"github.com/ardanlabs/service/app/domain/userapp"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI(log *logger.Logger, userBus *userbus.Business) *web.App {
	app := web.NewApp(log.Info, mid.Logger(log), mid.Error(log), mid.Panics())

	test.Routes(log, app)

	userapp.Routes(userBus, app)

	return app
}
