# [HTMX Rulez D00d](https://www.youtube.com/watch?v=1XUeiYieYK4)

Template Repository for creating [Go](https://go.dev/) / [Chi](https://github.com/go-chi/chi) / [Templ](https://templ.guide/) / [HTMX](https://htmx.org/) / [Tailwind](https://tailwindcss.com/) projects.

## How

Use this Repo Template to create your own repo. You will have to do a find and replace for `htmx-rules-dood` (with o's no zeros) with your go package name

## Install

Run `npm install` to install the Tailwind dependencies.

## Run

### Docker

This project uses Air and Docker Compose so you can compile Templ files and have live reloading without installing
extra dependencies on your local machine.

To start the project run

```sh
make up
```

This starts the server in a docker container that can be access at `localhost:8080`.

### Tailwind

To run the Tailwind css compiler, in a seperate terminal, run

```
make css
```

This will watch all `.templ` files for tailwind classes and rebuild the `dist/output.css` file.
