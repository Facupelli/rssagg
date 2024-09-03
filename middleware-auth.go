package main

import (
	"fmt"
	"net/http"

	"github.com/Facupelli/rssagg/internal/auth"
	"github.com/Facupelli/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
	
		user, err := apiCfg.DB.GetUserByApPIey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("Could not find user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
