package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerNotesGet(w http.ResponseWriter, r *http.Request, user User) {
	notes := []map[string]string{{"id": "1", "text": "Example note"}}
	respondWithJSON(w, http.StatusOK, notes)
}

func (cfg *apiConfig) handlerNotesCreate(w http.ResponseWriter, r *http.Request, user User) {
	var note map[string]string
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid note format")
		return
	}
	respondWithJSON(w, http.StatusCreated, note)
}
