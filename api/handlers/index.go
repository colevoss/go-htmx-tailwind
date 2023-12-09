package handlers

import (
	"net/http"
	"time"

	"htmx-rulez-dood/app"
	"htmx-rulez-dood/views"
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

func (i *IndexHandlers) GetStarted(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 1)
	comp := views.GetStarted("D00d")
	comp.Render(r.Context(), w)
}
