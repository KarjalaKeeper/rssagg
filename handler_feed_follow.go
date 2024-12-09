package main

import (
	"MainProject/internal/database"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// обработчик создания пользователя
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	parms := parameters{}
	err := decoder.Decode(&parms)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    parms.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't create feed follow:%v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollow(feedFollow))
}

// обработчик для просмотра подписок пользователя
func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't get feed follows:%v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

// удаление подписки

func (apiCfg *apiConfig) handlerDeleteGFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn`t parse feed: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn`t delete feed follow:%v", err))
	}
	respondWithJSON(w, 200, struct{}{})
}
