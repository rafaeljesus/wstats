package net

import (
	"io/ioutil"
	"net"
)

func ListenAndServeTCP(port string, rc chan string) {
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

		go handleConnection(conn, rc)
	}
}

func handleConnection(conn net.Conn, rc chan string) {
	defer conn.Close()
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		conn.Write([]byte("HTTP/1.0 Status 503 ServiceUnavailable\r\n\r\n"))
		return
	}

	rc <- string(buf)
	conn.Write([]byte("HTTP/1.0 Status 200 OK\r\n\r\n"))
}
