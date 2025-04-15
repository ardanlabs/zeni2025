package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

type App struct {
	*http.ServeMux
	mw []MidFunc
}

func NewApp(mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc IS NOT OWN IMPLEMENTATION
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.New())

		resp := handlerFunc(ctx, r)

		Respond(ctx, w, resp)

		// POST-PROCESSING
	}

	a.ServeMux.HandleFunc(pattern, h)
}
