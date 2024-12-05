package main

import (
	"MainProject/internal/database"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// обработчик создания пользователя
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	parms := parameters{}
	err := decoder.Decode(&parms)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parms.Name,
		Url:       parms.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't create feed:%v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

// обработчик получения каналов
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
