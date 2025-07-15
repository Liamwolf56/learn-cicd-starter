package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithJSON writes a JSON response with the given status code
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(data); err != nil {
		log.Printf("Error writing response body: %v", err)
	}
}

// respondWithError writes an error response in JSON format
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
