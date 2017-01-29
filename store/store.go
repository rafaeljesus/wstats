package store

import (
	"errors"
	"sort"
	"sync"
)

var ErrKeyNotExist = errors.New("key does not exist")

type Store struct {
	sync.RWMutex
	TotalCount int
	Wmap       map[string]Memkv
	Lmap       map[string]Memkv
}

func NewStore() *Store {
	return &Store{
		Wmap: make(map[string]Memkv),
		Lmap: make(map[string]Memkv),
	}
}

func (s *Store) IncWmap(key string, n int) (int, error) {
	s.Lock()
	defer s.Unlock()

	v, found := s.Wmap[key]
	if !found {
		v.Value = 0
	}

	nv := v.Value + n
	s.TotalCount++

	s.Wmap[key] = Memkv{key, nv}
	return nv, nil
}

func (s *Store) IncLmap(key string, n int) (int, error) {
	s.Lock()
	defer s.Unlock()

	v, found := s.Lmap[key]
	if !found {
		v.Value = 0
	}

	nv := v.Value + n
	s.TotalCount++

	s.Lmap[key] = Memkv{key, nv}
	return nv, nil
}

func (s *Store) Getw(key string) (Memkv, error) {
	s.Lock()
	defer s.Unlock()

	v, found := s.Wmap[key]
	if !found {
		return v, ErrKeyNotExist
	}

	return v, nil
}

func (s *Store) Getl(key string) (Memkv, error) {
	s.Lock()
	defer s.Unlock()

	v, found := s.Lmap[key]
	if !found {
		return v, ErrKeyNotExist
	}

	return v, nil
}

func (s *Store) SortByWords() ([]string, int) {
	s.Lock()
	defer s.Unlock()

	sm := new(SortMap)
	count := 0
	sm.mkv = s.Wmap
	sm.sorted = make([]string, len(s.Wmap))
	i := 0
	for key := range s.Wmap {
		count += s.Wmap[key].Value
		sm.sorted[i] = key
		i++
	}

	sort.Sort(sm)

	return sm.sorted, count
}

func (s *Store) SortByLetters() []string {
	sm := new(SortMap)
	sm.mkv = s.Lmap
	sm.sorted = make([]string, len(s.Lmap))
	i := 0
	for key := range s.Lmap {
		sm.sorted[i] = key
		i++
	}

	sort.Sort(sm)

	return sm.sorted
}

func (s *Store) Count() int {
	return s.TotalCount
}
