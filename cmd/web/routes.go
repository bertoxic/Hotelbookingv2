package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/bertoxic/bert/internal/config"
	"github.com/bertoxic/bert/internal/handlers"

)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionsLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/example", handlers.Repo.Example)
	mux.Get("/general-quaters", handlers.Repo.General)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/Book-room", handlers.Repo.BookRoom)

	mux.Get("/user-login", handlers.Repo.ShowLogin)
	mux.Post("/user-login", handlers.Repo.PostShowLogin)
	mux.Get("/user-logout", handlers.Repo.Logout)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	mux.Route("/admin",func(mux chi.Router) {
		//mux.Use(Auth)
		mux.Get("/dashboard",handlers.Repo.AdminDashboard)
		mux.Get("/reservations-new",handlers.Repo.AdminNewReservations)
		mux.Get("/reservations-all",handlers.Repo.AdminAllReservations)
		mux.Get("/reservation-calender",handlers.Repo.AdminReservationCalender)
		mux.Post("/reservation-calender",handlers.Repo.AdminPostReservationCalender)
		mux.Get("/process-reservations/{src}/{id}/do",handlers.Repo.AdminProcessReservation)
		mux.Get("/delete-reservations/{src}/{id}/do",handlers.Repo.AdminDeleteReservation)

		mux.Get("/reservation/{src}/{id}/show",handlers.Repo.AdminShowReservation)
		mux.Post("/reservation/{src}/{id}",handlers.Repo.AdminModifyReservation)
		
	})

	return mux

}
