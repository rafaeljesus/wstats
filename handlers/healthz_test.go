package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHealthzIndex(t *testing.T) {
	env := &Env{}

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/healthz", nil)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthz", env.Healthz)
	mux.ServeHTTP(res, req)

	response := make(map[string]string)
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Errorf("Expected to decode response json %s", err)
	}

	if response["status"] != "up" {
		t.Errorf("Expected status to equal %s", response["status"])
	}

	if res.Code != http.StatusOK {
		t.Error("Expected status %s to be equal %s", strconv.Itoa(res.Code), http.StatusOK)
	}
}
