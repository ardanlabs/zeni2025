package web

import (
	"context"
	"net/http"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

// HandleFunc IS NOT OWN IMPLEMENTATION
func (app *App) HandleFunc(pattern string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		// PRE-PROCESSING

		err := handler(r.Context(), w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// POST-PROCESSING
	}

	app.ServeMux.HandleFunc(pattern, h)
}
