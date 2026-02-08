package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	
	expected := "Hello, E2E Test!"
	if w.Body.String() != expected {
		t.Errorf("Expected %q, got %q", expected, w.Body.String())
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	healthHandler(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	
	expected := "{\"status\":\"healthy\",\"version\":\"1.0.0\"}"
	if w.Body.String() != expected {
		t.Errorf("Expected %q, got %q", expected, w.Body.String())
	}
}