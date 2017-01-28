package store

type Repo interface {
	IncWmap(key string, n int) (int, error)
	IncLmap(key string, n int) (int, error)
	Getw(key string) (Memkv, error)
	Getl(key string) (Memkv, error)
	SortedByWords() ([]string, int)
	SortedByLetters() []string
	Count() int
}
