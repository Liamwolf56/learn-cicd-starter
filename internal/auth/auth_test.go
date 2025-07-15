package auth

import (
    "os"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Set the expected API key for testing
    os.Setenv("API_KEY", "my-secret-key-123")

    apiKey, err := GetAPIKey()
    if err != nil {
        t.Errorf("did not expect error, got: %v", err)
    }

    if apiKey != "my-secret-key-123" {
        t.Errorf("expected 'my-secret-key-123', got %q", apiKey)
    }
}

