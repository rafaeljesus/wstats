package store

type Sort struct {
	mkv    map[string]Memkv
	sorted []string
}

func (s *Sort) Len() int {
	return len(s.mkv)
}

func (s *Sort) Less(i, j int) bool {
	if s.mkv[s.sorted[i]] == s.mkv[s.sorted[j]] {
		return s.sorted[i] < s.sorted[j]
	} else {
		return s.mkv[s.sorted[i]].Value > s.mkv[s.sorted[j]].Value
	}
}

func (s *Sort) Swap(i, j int) {
	s.sorted[i], s.sorted[j] = s.sorted[j], s.sorted[i]
}
