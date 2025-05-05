package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahsan0098/gofirst/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request,user database.User) {

	type parameters struct {
		Name   string `json:"name"`
		Url    string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("failed to create feed: %v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, feed)
}

func (apiCfg *apiConfig) HandleGetFeed(w http.ResponseWriter, r *http.Request) {

	feedId := r.URL.Query().Get("feed")
	if feedId == "" {
		respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("Feed id not found in request body: %v", feedId))
		return
	}

	idd, err := uuid.Parse(feedId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid user_id: %v", err))
		return
	}

	feed, err := apiCfg.DB.GetFeed(r.Context(), idd)
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("failed to get feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

func (apiCfg *apiConfig) HandleGetUserFeeds(w http.ResponseWriter, r *http.Request, user database.User) {

	
	feeds, err := apiCfg.DB.GetUserFeeds(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("failed to get feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, feeds)
}