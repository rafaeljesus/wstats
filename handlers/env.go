package handlers

import (
	"github.com/rafaeljesus/wstats/store"
)

type Env struct {
	ReceiveChannel chan string
	Repo           store.Repo
}

func NewEnv(repo store.Repo) *Env {
	return &Env{
		Repo:           repo,
		ReceiveChannel: make(chan string),
	}
}
