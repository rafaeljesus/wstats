package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/rafaeljesus/wstats/handlers"
	"github.com/rafaeljesus/wstats/net"
	"github.com/rafaeljesus/wstats/store"
	"runtime"
)

func main() {
	tcpPort := flag.String("tcp_port", "5555", "Set tcp port")
	flag.Parse()

	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)

	store := store.NewStore()
	env := handlers.NewEnv(store)

	log.WithField("tcp_port", tcpPort).Info("starting tcp server")

	net.ListenAndServeTCP(":"+*tcpPort, env.ReceiveChannel)
}
