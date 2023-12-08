package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/colevoss/go-htmx-tailwind/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		comp := views.Index()
		comp.Render(r.Context(), w)
	})

	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		comp := views.Hello(name)
		comp.Render(r.Context(), w)
	})

	FileServer(r)

	http.ListenAndServe(":3000", r)
}

var workDir, _ = os.Getwd()
var dir = http.Dir(filepath.Join(workDir, "dist"))

func FileServer(r *chi.Mux) {
	r.Get("/dist/*", func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(dir))
		fs.ServeHTTP(w, r)
	})
}
