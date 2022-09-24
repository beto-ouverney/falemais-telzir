package main

import (
	_ "github.com/beto-ouverney/falemais-telzir/docs"
	"github.com/beto-ouverney/falemais-telzir/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

// @title           FaleMais Telzir API
// @version         1.0
// @description     This is a sample server for Telzir.

// @contact.name   API Support
// @contact.url    https://www.linkedin.com/in/beto-ouverney-paz/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1/dddcost

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1/dddcost", func(r chi.Router) {
		r.Get("/", handler.GetAllDDDCodes)
		r.Post("/", handler.GetCost)
	})

	router.Mount("/api/v1/dddcost/swagger", httpSwagger.WrapHandler)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	log.Println("Server running on port: ", port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Printf("Failed to launch api server:%+v\n", err)
	}
}
