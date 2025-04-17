package userapp

import (
	"github.com/ardanlabs/service/business/domain/userbus"
	"github.com/ardanlabs/service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(userBus *userbus.Business, app *web.App) {
	const version = "v1"

	api := newApp(userBus)

	app.HandleFunc("GET /users", api.query)
	app.HandleFunc("GET /users/{user_id}", api.queryByID)
	app.HandleFunc("POST /users", api.create)
	app.HandleFunc("PUT /users/role/{user_id}", api.updateRole)
	app.HandleFunc("PUT /users/{user_id}", api.update)
	app.HandleFunc("DELETE /users/{user_id}", api.delete)
}
