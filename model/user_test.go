package model

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestUser_Insert(t *testing.T) {
	initDB()
	defer mysql.Close()

	u := &User{
		Name:     "jdxj",
		Password: "jdxj",
	}
	if err := u.Insert(); err != nil {
		t.Fatalf("%s\n", err)

	}
}

type foo struct {
	n int
	sync.Mutex
}

func TestCopyMutex(t *testing.T) {
	f := foo{n: 17}

	f.Lock()
	log.Println("g1: lock foo ok") // 1

	// 在mutex首次使用后复制其值
	go func(f foo) {
		for {
			log.Println("g3: try to lock foo...") // 4 一直阻塞
			f.Lock()
			log.Println("g3: lock foo ok")
			time.Sleep(5 * time.Second)
			f.Unlock()
			log.Println("g3: unlock foo ok")
		}
	}(f)

	time.Sleep(2 * time.Second)
	f.Unlock()
	log.Println("g1: unlock foo ok")
	time.Sleep(10 * time.Second)
	log.Println("g1: unlock foo ok 2")
}
