package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header missing")
	}

	if !strings.HasPrefix(authHeader, "ApiKey ") {
		return "", errors.New("Authorization header format must be 'ApiKey <key>'")
	}

	return strings.TrimPrefix(authHeader, "ApiKey "), nil
}
