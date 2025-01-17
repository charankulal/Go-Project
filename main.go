package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is Not found")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router:= chi.NewRouter()

	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/err",handlerErr)

	router.Mount("/v1",v1Router)

	serv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server Starting on port %v",portString)

	err := serv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
