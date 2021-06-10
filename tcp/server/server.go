package server

import (
	"log"
	"net"
)

func Start() {
	l, err := net.Listen("tcp", "127.0.0.1:49152")
	if err != nil {
		log.Fatalf("Listen: %s\n", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Accept: %s\n", err)
			continue
		}

		go func(c net.Conn) {
			buf := make([]byte, 1024)
			c.Read(buf)
		}(conn)
	}
}
