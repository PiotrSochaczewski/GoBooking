package main

import (
	"net/http"

	"github.com/PiotrSochaczewski/GoBooking/pkg/config"
	"github.com/PiotrSochaczewski/GoBooking/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	//Get Rooms
	mux.Get("/modern-hanok", handlers.Repo.ModernHanok)
	mux.Get("/traditional-hanok", handlers.Repo.TraditionalHanok)
	mux.Get("/modern", handlers.Repo.Modern)

	//SearchAvailability GET and POST
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
