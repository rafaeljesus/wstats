package net

import (
	"github.com/rafaeljesus/wstats/worker"
	"io/ioutil"
	"net"
)

func ListenAndServeTCP(port string, rq chan worker.Request) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn, rq)
	}
}

func handleConnection(conn net.Conn, rq chan worker.Request) {
	defer conn.Close()
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		conn.Write([]byte("HTTP/1.0 Status 503 ServiceUnavailable\r\n\r\n"))
		return
	}

	request := worker.Request{Payload: string(buf)}
	rq <- request
	conn.Write([]byte("HTTP/1.0 Status 200 OK\r\n\r\n"))
}
