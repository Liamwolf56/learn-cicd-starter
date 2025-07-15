package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// ✅ Use your real API key in the test (as requested)
	const myAPIKey = "AIzaSyBhcC9jHtODH2Qkc-hGfsJDjYdpiIyXRk8"

	header := http.Header{}
	header.Set("Authorization", "ApiKey "+myAPIKey)

	apiKey, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if apiKey != myAPIKey {
		t.Errorf("expected API key %q, got %q", myAPIKey, apiKey)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	header := http.Header{}

	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected an error when Authorization header is missing")
	}
}

func TestGetAPIKey_InvalidPrefix(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "Bearer sometoken")

	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatal("expected an error when Authorization header has wrong prefix")
	}
}
