package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"user-service-api/routes"
	"user-service-api/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	storage.InitDB(dsn)

	r := chi.NewRouter()

	r = routes.RegisterRoutes()

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
