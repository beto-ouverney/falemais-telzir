package main

import (
	"github.com/beto-ouverney/falemais-telzir/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1/dddcost", func(r chi.Router) {
		r.Get("/", handler.GetAllDDDCodes)
		r.Post("/", handler.GetCost)
	})

	port := os.Getenv("PORT")
	log.Println("Server running on port " + port)
	http.ListenAndServe(port, router)
}
