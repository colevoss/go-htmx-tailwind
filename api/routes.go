package api

import "github.com/go-chi/chi/v5"

func (a *Api) Routes() {
	a.api.Route("/", func(r chi.Router) {
		r.Get("/", a.handlers.Index.Index)
		r.Get("/get-started", a.handlers.Index.GetStarted)
		r.Get("/{name}", a.handlers.Index.Hello)
	})

	a.api.Route("/dist", func(r chi.Router) {
		r.Get("/*", a.handlers.Files.Serve)
	})
}
