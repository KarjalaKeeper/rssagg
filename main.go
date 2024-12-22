package main

import (
	"MainProject/internal/database"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	feed, err := urlToFeed("https://techcrunch.com/feed/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)

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

	db := database.New(conn)
	//конфигурация API
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

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

	//подключаем handlerGetUser
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	//подключаем handlerCreateFeed
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	//подключаем handlerGetFeeds
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	//подключаем handlerCreateFeedFollow
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))

	//handlerGetFeedFollows
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))

	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteGFeedFollow))

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
