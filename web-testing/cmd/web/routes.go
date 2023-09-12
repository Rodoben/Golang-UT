package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)
	return mux
}
