package main

import (
	"os"

	"github.com/jdxj/share/tcp/client"
	"github.com/jdxj/share/tcp/server"
)

func main() {
	s := os.Args[1]
	switch s {
	case "server":
		server.Start()
	case "client":
		client.Dial()
	}
}
