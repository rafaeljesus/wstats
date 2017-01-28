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

func (s *Store) SortedByWords() ([]string, int) {
	s.Lock()
	defer s.Unlock()

	ws := new(Sort)
	count := 0
	ws.mkv = s.Wmap
	ws.sorted = make([]string, len(s.Wmap))
	i := 0
	for key, _ := range s.Wmap {
		count += s.Wmap[key].Value
		ws.sorted[i] = key
		i++
	}

	sort.Sort(ws)

	return ws.sorted, count
}

func (s *Store) SortedByLetters() []string {
	s.Lock()
	defer s.Unlock()

	ws := new(Sort)
	ws.mkv = s.Lmap
	ws.sorted = make([]string, len(s.Lmap))
	i := 0
	for key, _ := range s.Lmap {
		ws.sorted[i] = key
		i++
	}

	sort.Sort(ws)

	return ws.sorted
}

func (s *Store) Count() int {
	return s.TotalCount
}
