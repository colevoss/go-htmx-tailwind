package handlers

import (
	"net/http"

	"github.com/colevoss/go-htmx-tailwind/app"
	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Index *IndexHandlers
	Files *FilesHandler
}

func NewHandlers(app *app.App) *Handlers {
	route := NewRoute()

	return &Handlers{
		Index: NewIndexHandlers(app, route),
		Files: NewFilesHandler(app),
	}
}

type Route struct {
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) UrlParam(req *http.Request, param string) string {
	return chi.URLParam(req, param)
}
