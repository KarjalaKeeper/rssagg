package main

import (
	"MainProject/internal/auth"
	"MainProject/internal/database"
	"fmt"
	"net/http"
)

//создаем пользовательский тип для обработчиков аутентификации

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// прослойка, котора аутенцифицирует запрос, получает пользователя и вызывает следующий обработчик
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
