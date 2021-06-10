package client

import (
	"log"
	"net"
)

func Dial() {
	conn, err := net.Dial("tcp", "127.0.0.1:49152")
	if err != nil {
		log.Fatalf("Dial: %s\n", err)
	}

	err = conn.Close()
	if err != nil {
		log.Printf("Close: %s\n", err)
	} else {
		log.Printf("Close ok1")
	}
	err = conn.Close()
	if err != nil {
		log.Printf("Close: %s\n", err)
	} else {
		log.Printf("Close ok2")
	}
}
