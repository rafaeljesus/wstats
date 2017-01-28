package store

type Memkvs []Memkv

func (ks Memkvs) Len() int {
	return len(ks)
}

func (ks Memkvs) Less(i, j int) bool {
	return ks[i].Value < ks[j].Value
}

func (ks Memkvs) Swap(i, j int) {
	ks[i], ks[j] = ks[j], ks[i]
}
