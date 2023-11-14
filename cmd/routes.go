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

	r.Get("/", handler.Repo.Home)
	r.Get("/about", handler.Repo.About)
	r.Get("/generals-quarters", handler.Repo.Generals)
	r.Get("/majors-suite", handler.Repo.Majors)
	r.Get("/search-availability", handler.Repo.Availability)
	r.Post("/search-availability", handler.Repo.PostAvailability)
	r.Get("/contact", handler.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return r
}
