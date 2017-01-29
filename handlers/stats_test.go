package handlers

import (
	"encoding/json"
	"github.com/rafaeljesus/wstats/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatsIndex(t *testing.T) {
	store := store.NewStore()
	env := NewEnv(store, nil)

	payload := "lorem input, input, tok;"
	env.StatsCreate(payload)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/stats", nil)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/stats", env.StatsIndex)
	mux.ServeHTTP(res, req)

	response := &statsReponse{}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		t.Errorf("Expected to decode stats response json %s", err)
	}

	if response.Count != 4 {
		t.Errorf("Expected count to be eq %s", response.Count)
	}

	if response.Total != 4 {
		t.Errorf("Expected total count to be q %s", response.Total)
	}

	if response.TopWords[0] != "input" || response.TopWords[4] != "lorem" {
		t.Errorf("Expected top words to be ordered%s", response.TopWords)
	}

	if response.TopLetters[0] != "t" || response.TopLetters[4] != "i" {
		t.Errorf("Expected top letters to be ordered%s", response.TopLetters)
	}
}

func TestStatsCreate(t *testing.T) {
	store := store.NewStore()
	env := NewEnv(store, nil)

	payload := "lorem input, input, tok;"
	env.StatsCreate(payload)

	if store.Count() != 22 {
		t.Errorf("Expected score count to be %s", store.Count())
	}

	kv, err := store.Getw("lorem")
	if err != nil {
		t.Errorf("Expected get \"lorem\" word", err)
	}

	kv, err = store.Getw("input")
	if err != nil {
		t.Errorf("Expected get \"input\" word", err)
	}

	if kv.Value != 2 {
		t.Errorf("Expected kv value to be 2 got %s", kv.Value)
	}

	kv, err = store.Getl("t")
	if err != nil {
		t.Errorf("Expected get \"t\" letter", err)
	}

	if kv.Value != 3 {
		t.Errorf("Expected kv value to be 3 got %s", kv.Value)
	}
}
