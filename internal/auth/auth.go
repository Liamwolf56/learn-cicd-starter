package auth

import (
	"context"
	"errors"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	ApiKey string `json:"api_key"`
}

func GetAPIKey(headers http.Header) (string, error) {
	key := headers.Get("X-API-Key")
	if key == "" {
		return "", errors.New("API key missing")
	}
	return key, nil
}

func GetUserByAPIKey(ctx context.Context, key string) (User, error) {
	// Stub user lookup logic
	if key == "valid-key" {
		return User{ID: 1, Email: "user@example.com", Name: "John Doe", ApiKey: key}, nil
	}
	return User{}, errors.New("invalid API key")
}
