package handlers

import (
	"net/http"

	"github.com/colevoss/go-htmx-tailwind/app"
	"github.com/colevoss/go-htmx-tailwind/views"
)

type IndexHandlers struct {
	*Route
	app *app.App
}

func NewIndexHandlers(app *app.App, route *Route) *IndexHandlers {
	return &IndexHandlers{
		app:   app,
		Route: route,
	}
}

func (i *IndexHandlers) Index(w http.ResponseWriter, r *http.Request) {
	comp := views.Index()
	i.app.Log.Info("Howdy Index")
	comp.Render(r.Context(), w)
}

func (i *IndexHandlers) Hello(w http.ResponseWriter, r *http.Request) {
	// name := chi.URLParam(r, "name")
	name := i.UrlParam(r, "name")
	comp := views.Hello(name)
	comp.Render(r.Context(), w)
}
