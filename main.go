package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/rafaeljesus/wstats/handlers"
	"github.com/rafaeljesus/wstats/net"
	"github.com/rafaeljesus/wstats/store"
	"github.com/rafaeljesus/wstats/worker"
	"net/http"
	"runtime"
)

func main() {
	tcpPort := flag.String("tcp_port", "5555", "Set tcp port")
	httpPort := flag.String("http_port", "8080", "Set http port")
	flag.Parse()

	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	store := store.NewStore()
	mux := http.NewServeMux()

	env := handlers.NewEnv(store, mux)
	task := env.StatsCreate

	dispatcher := worker.NewDispatcher(numcpu)
	dispatcher.Run(task)

	log.WithField("tcp_port", *tcpPort).Info("starting tcp server")
	go net.ListenAndServeTCP(":"+*tcpPort, worker.RequestQueue)

	mux.HandleFunc("/v1/healthz", env.Healthz)
	mux.HandleFunc("/v1/stats", env.StatsIndex)

	log.WithField("http_port", *httpPort).Info("starting http server")
	http.ListenAndServe(":"+*httpPort, env)
}
