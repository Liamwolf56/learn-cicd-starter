package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerNotesGet(w http.ResponseWriter, r *http.Request, user User) {
	notes := []string{"Note 1", "Note 2", "Note 3"} // placeholder
	respondWithJSON(w, http.StatusOK, notes)
}

func (cfg *apiConfig) handlerNotesCreate(w http.ResponseWriter, r *http.Request, user User) {
	var note map[string]string
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid note")
		return
	}
	respondWithJSON(w, http.StatusCreated, note)
}
