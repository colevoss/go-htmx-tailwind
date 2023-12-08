package main

import (
	"log/slog"
	"os"

	"github.com/colevoss/go-htmx-tailwind/api"
	"github.com/colevoss/go-htmx-tailwind/app"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &app.App{
		Log: l,
	}

	api := api.NewApi(app)

	l.Info("Hello", "hi", "howare you")

	api.Run()
}
