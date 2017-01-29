package store

type SortMap struct {
	mkv    map[string]Memkv
	sorted []string
}

func (s *SortMap) Len() int {
	return len(s.mkv)
}

func (s *SortMap) Less(i, j int) bool {
	return s.mkv[s.sorted[i]].Value > s.mkv[s.sorted[j]].Value
}

func (s *SortMap) Swap(i, j int) {
	s.sorted[i], s.sorted[j] = s.sorted[j], s.sorted[i]
}
