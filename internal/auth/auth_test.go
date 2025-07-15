package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Create a fake header with the correct API key
    header := http.Header{}
    header.Set("Authorization", "ApiKey AIzaSyBhcC9jHtODH2Qkc-hGfsJDjYdpiIyXRk8")

    apiKey, err := GetAPIKey(header)
    if err != nil {
        t.Errorf("did not expect error, got: %v", err)
    }

    expected := "AIzaSyBhcC9jHtODH2Qkc-hGfsJDjYdpiIyXRk8"
    if apiKey != expected {
        t.Errorf("expected %q, got %q", expected, apiKey)
    }
}

