package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer) // to recover if there is a panic
	mux.Use(app.addIPToContext)
	mux.Use(app.Session.LoadAndSave) // to use the session in middleware, persisting data
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)
	mux.Get("/user/profile", app.Profile)
	fileServer := http.FileServer(http.Dir("/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
