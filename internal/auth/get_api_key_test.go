package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKeyForWhenNoAuthHeaderIsProvided(t *testing.T) {
	authHeader := http.Header{}
	_, err := GetAPIKey(authHeader)

	if err == nil {
		t.Fatalf("Expected error when no auth header is provided, got nil.")
	}
}

func TestGetApiKeyForWhenAuthHeaderIsMalformed(t *testing.T) {
	authHeader := http.Header{"Authorization": {"Bearer something"}}
	_, err := GetAPIKey(authHeader)

	if err == nil {
		t.Fatalf("Expected error when auth header is malformed, got nil.")
	}
}

func TestGetApiKeyForWhenNoApiKeyIsProvided(t *testing.T) {
	authHeader := http.Header{"Authorization": {"ApiKey "}}
	_, err := GetAPIKey(authHeader)

	if err == nil {
		t.Fatalf("Expected error when no API key is provided, got nil.")
	}
}
