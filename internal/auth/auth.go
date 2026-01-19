package auth

import (
	"errors"
	"net/http"
)

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	v := headers.Get("Authorization")
	if v == "" {
		return "", errors.New("no authentication info found")
	}
	return v, nil
}
