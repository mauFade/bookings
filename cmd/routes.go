package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mauFade/bookings/pkg/config"
	"github.com/mauFade/bookings/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(NoSurf)
	r.Use(SessionLoad)

	r.Get("/", http.HandlerFunc(handler.Repo.Home))
	r.Get("/about", http.HandlerFunc(handler.Repo.About))

	return r
}
