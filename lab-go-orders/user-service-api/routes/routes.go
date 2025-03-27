package routes

import (
	"user-service-api/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/register", handlers.RegisterUser)
	return r
}
