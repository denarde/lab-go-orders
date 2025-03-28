package routes

import (
	"user-service-api/handlers"
	"user-service-api/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.RequestID)

	r.Post("/register", handlers.RegisterUser)
	return r
}
