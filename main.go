package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Загружаем переменные окружения из .env
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found it the environment")
	}
	// Создаём маршрутизатор
	router := chi.NewRouter()

	// Настраиваем CORS доменных запросов
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// // Создаём маршруты версии API
	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/ready", handlerReadiness)
	// Подключаем маршруты версии API к основному маршрутизатору
	router.Mount("/v1", v1Router)

	// Создаём HTTP-сервер
	srv := &http.Server{
		Handler: router, //назначаем маршрутизатор обработчиком запросов
		Addr:    ":" + portString,
	}
	log.Printf("Server running on port %v", portString) // Запускаем сервер
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
