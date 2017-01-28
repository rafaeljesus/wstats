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

func TestCount(t *testing.T) {
	s := NewStore()
	_, _ = s.IncWmap("key", 1)
	_, _ = s.IncLmap("key", 1)
	if s.Count != 2 {
		t.Errorf("Expected 2 equal total count ", s.Count)
	}
}
