package store

import (
	"testing"
)

func TestIncWmap(t *testing.T) {
	s := NewStore()
	n, err := s.IncWmap("key", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment word map", err)
	}

	n, err = s.IncWmap("key", 1)
	if err != nil || n != 2 {
		t.Errorf("Expected to increment word map", err)
	}
}

func TestIncLmap(t *testing.T) {
	s := NewStore()
	n, err := s.IncLmap("key", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment word map", err)
	}

	n, err = s.IncLmap("key", 1)
	if err != nil || n != 2 {
		t.Errorf("Expected to increment word map", err)
	}
}

func TestGetm(t *testing.T) {
	s := NewStore()
	n, err := s.IncWmap("key", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment word map", err)
	}

	_, err = s.Getw("not_exist")
	if err == nil {
		t.Errorf("Expected to not get a not existing word", err)
	}

	kv, err := s.Getw("key")
	if err != nil {
		t.Errorf("Expected get word", err)
	}

	if kv.Value != 1 {
		t.Errorf("Expected kv value to be 1 got %s", kv.Value)
	}
}

func TestCount(t *testing.T) {
	s := NewStore()
	_, _ = s.IncWmap("key", 1)
	_, _ = s.IncLmap("key", 1)
	if s.Count() != 2 {
		t.Errorf("Expected 2 equal total count ", s.Count())
	}
}
