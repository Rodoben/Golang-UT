package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *application) Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCors)
	mux.Post("/authentication", app.authenticate)
	mux.Post("/refresh-token", app.refreshToken)

	mux.Route("/users", func(mux chi.Router) {
		mux.Use(app.authRequired)
		mux.Get("/", app.allUser)
		mux.Get("/{userID}", app.getOneUser)
		mux.Put("/", app.insertUser)
		mux.Patch("/", app.updateUser)
		mux.Delete("/{userID}", app.deleteUser)
	})
	return mux
}
