package main

import (
	"context"
	"log"
	"os"
	"time"

	v6 "github.com/go-redis/redis"
	v8 "github.com/go-redis/redis/v8"
)

func main() {
	ver := os.Args[1]
	addr := os.Args[2]

	switch ver {
	case "v6":
		connectClusterV6(addr)
	case "v8":
		connectClusterV8(addr)
	}
}

func connectClusterV8(addr string) {
	addrs := []string{addr}
	opt := &v8.ClusterOptions{
		Addrs:    addrs,
		ReadOnly: true,
	}
	cc := v8.NewClusterClient(opt)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	for {
		err := cc.Ping(ctx).Err()
		if err != nil {
			log.Printf("Ping: %s\n", err)
		} else {
			log.Printf("ok\n")
		}

		select {
		case <-ticker.C:
		}
	}
}

func connectClusterV6(addr string) {
	addrs := []string{addr}
	opt := &v6.ClusterOptions{
		Addrs:    addrs,
		ReadOnly: true,
	}
	cc := v6.NewClusterClient(opt)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		err := cc.Ping().Err()
		if err != nil {
			log.Printf("Ping: %s\n", err)
		} else {
			log.Printf("ok\n")
		}

		select {
		case <-ticker.C:
		}
	}
}
