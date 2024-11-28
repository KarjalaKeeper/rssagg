package main

import (
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found it the environment")
	}
	routrer := chi.NewRouter()

	srv := &http.Server{
		Handler: routrer, //обработчик
		Addr:    ":" + portString,
	}
	log.Printf("Server running on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
