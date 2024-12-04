package main

import (
	"MainProject/internal/database"
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Загружаем переменные окружения из .env
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found it the environment")
	}
	//подключение к БД
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found it the environment")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	//конфигурация API

	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)                  //endpoint error
	v1Router.Post("/users", apiCfg.handlerCreateUser) //подключение обработчика пользователя
	// Подключаем маршруты версии API к основному маршрутизатору
	router.Mount("/v1", v1Router)

	// Создаём HTTP-сервер
	srv := &http.Server{
		Handler: router, //назначаем маршрутизатор обработчиком запросов
		Addr:    ":" + portString,
	}
	log.Printf("Server running on port %v", portString) // Запускаем сервер
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
