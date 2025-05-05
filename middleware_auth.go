package main

import (
	"fmt"
	"net/http"

	"github.com/ahsan0098/gofirst/internal/auth"
	"github.com/ahsan0098/gofirst/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(next authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("failed to authorize user: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("failed to get user: %v", err))
			return
		}

		next(w, r, user)
	}
}
