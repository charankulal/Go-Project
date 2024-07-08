package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is Not found")
	}

	router := chi.NewRouter()

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
