package main

import (
	"net/http"

	"github.com/Totiruzi/reservation/pkg/config"
	"github.com/Totiruzi/reservation/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(PrintToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contct", handlers.Repo.Contact)
	mux.Get("/esimi", handlers.Repo.Esimi)
	mux.Get("/lavinda", handlers.Repo.Lavinda)
	mux.Get("/omoche", handlers.Repo.Omoche)
	mux.Get("/reservation", handlers.Repo.Reservastion)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
