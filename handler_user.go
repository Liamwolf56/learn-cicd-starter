package main

import (
	"net/http"
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user auth.User) {
	respondWithJSON(w, http.StatusOK, user)
}
