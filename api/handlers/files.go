package handlers

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"htmx-rulez-dood/app"
)

type FilesHandler struct {
	dir http.Dir
	fs  http.Handler
}

func NewFilesHandler(app *app.App) *FilesHandler {
	workDir, _ := os.Getwd()
	dir := http.Dir(filepath.Join(workDir, "dist"))

	mime.AddExtensionType(".js", "text/javascript")
	mime.AddExtensionType(".mjs", "text/javascript")

	fs := http.FileServer(dir)

	t := mime.TypeByExtension(".mjs")
	fmt.Printf("t: %v\n", t)

	return &FilesHandler{
		dir,
		fs,
	}
}

func (f *FilesHandler) Serve(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/dist", f.fs)

	fs.ServeHTTP(w, r)
}
