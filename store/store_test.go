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

func TestGetl(t *testing.T) {
	s := NewStore()
	n, err := s.IncLmap("a", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment letter map", err)
	}

	_, err = s.Getl("z")
	if err == nil {
		t.Errorf("Expected to not get a not existing letter", err)
	}

	kv, err := s.Getl("a")
	if err != nil {
		t.Errorf("Expected get letter", err)
	}

	if kv.Value != 1 {
		t.Errorf("Expected kv value to be 1 got %s", kv.Value)
	}
}

func TestSortByWords(t *testing.T) {
	s := NewStore()
	_, err := s.IncWmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment word map", err)
	}

	_, err = s.IncWmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment word map", err)
	}

	_, err = s.IncWmap("bar", 1)
	if err != nil {
		t.Errorf("Expected to increment word map", err)
	}

	sorted, count := s.SortedByWords()

	if sorted[0] != "foo" && sorted[1] != "bar" {
		t.Errorf("Expected sort to ordered by top words%s", sorted)
	}

	if count != 3 {
		t.Errorf("Expected count value to be %s", count)
	}
}

func TestSortByLetters(t *testing.T) {
	s := NewStore()
	_, err := s.IncLmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment letter map", err)
	}

	_, err = s.IncLmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment letter map", err)
	}

	_, err = s.IncLmap("bar", 1)
	if err != nil {
		t.Errorf("Expected to increment letter map", err)
	}

	sorted := s.SortedByLetters()

	if sorted[0] != "foo" && sorted[1] != "bar" {
		t.Errorf("Expected sort to ordered by top letters%s", sorted)
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
