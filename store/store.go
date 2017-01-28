package store

import (
	"sync"
)

type Repo interface {
	IncWmap(key string, n int) (int, error)
	IncLmap(key string, n int) (int, error)
	Count() int
}

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

func (s *Store) Count() int {
	return s.TotalCount
}
