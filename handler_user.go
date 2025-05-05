package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahsan0098/gofirst/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid request body: %v", err))
		return
	}

	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("failed to create user: %v", err))
		return
	}
	respondWithJSON(w, http.StatusCreated, newUser)
}

func (apiCfg *apiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, http.StatusOK, user)
}
