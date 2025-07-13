package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Case 1: Missing Authorization header
	headers1 := http.Header{}
	key, err := GetAPIKey(headers1)
	if err == nil {
		t.Errorf("expected error when Authorization header is missing")
	}
	if key != "" {
		t.Errorf("expected empty key when header is missing, got %q", key)
	}

	// Case 2: Malformed Authorization header
	headers2 := http.Header{}
	headers2.Set("Authorization", "Bearer") // missing actual key
	key, err = GetAPIKey(headers2)
	if err == nil {
		t.Errorf("expected error when Authorization header is malformed")
	}
	if key != "" {
		t.Errorf("expected empty key for malformed header, got %q", key)
	}

	// Case 3: Valid Authorization header
	headers3 := http.Header{}
	headers3.Set("Authorization", "Bearer AIzaSyA-X4VUXIF7MMpfgsPKShpsOWWvbeQ6pwQ") // Replace if needed
	key, err = GetAPIKey(headers3)
	if err != nil {
		t.Errorf("did not expect error, got: %v", err)
	}
	if key != "AIzaSyA-X4VUXIF7MMpfgsPKShpsOWWvbeQ6pwQ" {
		t.Errorf("expected 'AIzaSyA-X4VUXIF7MMpfgsPKShpsOWWvbeQ6pwQ', got %q", key)
	}
}

