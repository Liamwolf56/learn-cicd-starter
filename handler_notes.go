package main

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handlerNotesGet(w http.ResponseWriter, r *http.Request, user auth.User) {
	respondWithJSON(w, http.StatusOK, []string{"note 1", "note 2"})
}

func (cfg *apiConfig) handlerNotesCreate(w http.ResponseWriter, r *http.Request, user auth.User) {
	respondWithJSON(w, http.StatusCreated, map[string]string{"status": "created"})
}
