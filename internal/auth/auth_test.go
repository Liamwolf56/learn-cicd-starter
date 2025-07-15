package auth

import (
    "os"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Set the API key environment variable for testing
    os.Setenv("API_KEY", "AIzaSyBhcC9jHtODH2Qkc-hGfsJDjYdpiIyXRk8")

    apiKey, err := GetAPIKey()
    if err != nil {
        t.Errorf("did not expect error, got: %v", err)
    }

    if apiKey != "AIzaSyBhcC9jHtODH2Qkc-hGfsJDjYdpiIyXRk8" {
        t.Errorf("expected 'AIzaSyBhcC9jHtODH2Qkc-hGfsJDjYdpiIyXRk8', got %q", apiKey)
    }
}

