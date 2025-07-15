package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	var input map[string]string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	respondWithJSON(w, http.StatusCreated, input)
}

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user User) {
	respondWithJSON(w, http.StatusOK, user)
}
