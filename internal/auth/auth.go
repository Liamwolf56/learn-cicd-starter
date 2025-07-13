package auth

import (
    "errors"
    "net/http"
    "strings"
)

func GetAPIKey(headers http.Header) (string, error) {
    // Get the Authorization header
    authHeader := headers.Get("Authorization")
    if authHeader == "" {
        return "", errors.New("no authorization header included")
    }

    // Should be "Bearer <key>"
    parts := strings.SplitN(authHeader, " ", 2)
    if len(parts) != 2 {
        return "", errors.New("malformed authorization header")
    }

    if strings.ToLower(parts[0]) != "bearer" {
        return "", errors.New("authorization header must start with Bearer")
    }

    key := parts[1]
    if key == "" {
        return "", errors.New("authorization header missing API key")
    }

    return "", nil
}

