package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/colevoss/go-htmx-tailwind/app"
	"github.com/go-chi/chi/v5"
)

type FilesHandler struct {
	dir http.Dir
}

func NewFilesHandler(app *app.App) *FilesHandler {
	workDir, _ := os.Getwd()
	dir := http.Dir(filepath.Join(workDir, "dist"))

	return &FilesHandler{
		dir,
	}
}

func (f *FilesHandler) Serve(w http.ResponseWriter, r *http.Request) {
	rctx := chi.RouteContext(r.Context())
	pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
	fs := http.StripPrefix(pathPrefix, http.FileServer(f.dir))
	fs.ServeHTTP(w, r)
}
