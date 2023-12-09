package api

import (
	"net/http"

	"htmx-rulez-dood/api/handlers"
	"htmx-rulez-dood/app"

	"github.com/go-chi/chi/v5"
)

type Api struct {
	api      *chi.Mux
	app      *app.App
	handlers *handlers.Handlers
}

func NewApi(app *app.App) *Api {
	return &Api{
		api:      chi.NewRouter(),
		app:      app,
		handlers: handlers.NewHandlers(app),
	}
}

func (a *Api) Run() {
	a.Routes()
	http.ListenAndServe(":3000", a.api)
}
