package handlers

import (
	"github.com/rafaeljesus/wstats/store"
	"testing"
)

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
