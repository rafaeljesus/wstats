package store

import (
	"strconv"
	"testing"
)

func TestIncWmap(t *testing.T) {
	s := NewStore()
	n, err := s.IncWmap("key", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment word map %s", err)
	}

	n, err = s.IncWmap("key", 1)
	if err != nil || n != 2 {
		t.Errorf("Expected to increment word map %s", err)
	}
}

func TestIncLmap(t *testing.T) {
	s := NewStore()
	n, err := s.IncLmap("key", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment word map %s", err)
	}

	n, err = s.IncLmap("key", 1)
	if err != nil || n != 2 {
		t.Errorf("Expected to increment word map %s", err)
	}
}

func TestGetm(t *testing.T) {
	s := NewStore()
	n, err := s.IncWmap("key", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment word map %s", err)
	}

	_, err = s.Getw("not_exist")
	if err == nil {
		t.Errorf("Expected to not get a not existing word %s", err)
	}

	kv, err := s.Getw("key")
	if err != nil {
		t.Errorf("Expected get word %s", err)
	}

	if kv.Value != 1 {
		t.Errorf("Expected kv value to be 1 got %s", strconv.Itoa(kv.Value))
	}
}

func TestGetl(t *testing.T) {
	s := NewStore()
	n, err := s.IncLmap("a", 1)
	if err != nil || n != 1 {
		t.Errorf("Expected to increment letter map %s", err)
	}

	_, err = s.Getl("z")
	if err == nil {
		t.Errorf("Expected to not get a not existing letter %s", err)
	}

	kv, err := s.Getl("a")
	if err != nil {
		t.Errorf("Expected get letter %s", err)
	}

	if kv.Value != 1 {
		t.Errorf("Expected kv value to be 1 got %s", strconv.Itoa(kv.Value))
	}
}

func TestSortByWords(t *testing.T) {
	s := NewStore()
	_, err := s.IncWmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment word map %s", err)
	}

	_, err = s.IncWmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment word map %s", err)
	}

	_, err = s.IncWmap("bar", 1)
	if err != nil {
		t.Errorf("Expected to increment word map %s", err)
	}

	sorted, count := s.SortByWords()

	if sorted[0] != "foo" && sorted[1] != "bar" {
		t.Errorf("Expected sort to ordered by top words %s", sorted)
	}

	if count != 3 {
		t.Errorf("Expected count value to be %s", strconv.Itoa(count))
	}
}

func TestSortByLetters(t *testing.T) {
	s := NewStore()
	_, err := s.IncLmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment letter map %s", err)
	}

	_, err = s.IncLmap("foo", 1)
	if err != nil {
		t.Errorf("Expected to increment letter map %s", err)
	}

	_, err = s.IncLmap("bar", 1)
	if err != nil {
		t.Errorf("Expected to increment letter map %s", err)
	}

	sorted := s.SortByLetters()

	if sorted[0] != "foo" && sorted[1] != "bar" {
		t.Errorf("Expected sort to ordered by top letters %s", sorted)
	}
}

func TestCount(t *testing.T) {
	s := NewStore()
	_, _ = s.IncWmap("key", 1)
	_, _ = s.IncLmap("key", 1)
	if s.Count() != 2 {
		t.Errorf("Expected 2 equal total count %s", strconv.Itoa(s.Count()))
	}
}
