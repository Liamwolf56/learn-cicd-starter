package main

import (
	"net/http"
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, auth.User)

func (cfg *apiConfig) middlewareAuth(h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "missing or malformed API key")
			return
		}
		user, err := auth.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "invalid API key")
			return
		}
		h(w, r, user)
	}
}
