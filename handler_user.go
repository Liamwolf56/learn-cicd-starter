package main

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user auth.User) {
	respondWithJSON(w, http.StatusOK, user)
}
