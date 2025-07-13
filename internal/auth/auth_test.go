package auth

import (
    "os"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Backup existing env variable if any
    original := os.Getenv("API_KEY")
    defer os.Setenv("API_KEY", original) // Restore after test

    // Test case: environment variable not set
    os.Unsetenv("API_KEY")
    key := GetAPIKey()
    if key != "" {
        t.Errorf("expected empty string when API_KEY not set, got %q", key)
    }

    // Test case: environment variable set
    expected := "secret123"
    os.Setenv("API_KEY", expected)
    key = GetAPIKey()
    if key != expected {
        t.Errorf("expected %q, got %q", expected, key)
    }
}
