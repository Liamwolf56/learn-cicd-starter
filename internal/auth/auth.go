package auth

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrNoAuthHeader        = errors.New("no authorization header included")
	ErrMalformedAuthHeader = errors.New("malformed authorization header")
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrMalformedAuthHeader
	}

	return parts[1], nil
}
