package handlers

import (
	"github.com/rafaeljesus/wstats/store"
	"net/http"
)

type Env struct {
	Mux  *http.ServeMux
	Repo store.Repo
}

func NewEnv(repo store.Repo, mux *http.ServeMux) *Env {
	return &Env{
		Mux:  mux,
		Repo: repo,
	}
}

func (e *Env) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.Mux.ServeHTTP(w, r)
}
