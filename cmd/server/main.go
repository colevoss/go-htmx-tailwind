package main

import (
	"log/slog"
	"os"

	"htmx-rulez-dood/api"
	"htmx-rulez-dood/app"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &app.App{
		Log: l,
	}

	api := api.NewApi(app)

	l.Info("HTMX", "Rulez", "D00d")

	api.Run()
}
