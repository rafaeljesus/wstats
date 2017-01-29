package store

type Repo interface {
	IncWmap(key string, n int) (int, error)
	IncLmap(key string, n int) (int, error)
	Getw(key string) (Memkv, error)
	Getl(key string) (Memkv, error)
	SortByWords() ([]string, int)
	SortByLetters() []string
	Count() int
}
