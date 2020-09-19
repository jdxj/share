package model

import (
	"fmt"
	"testing"
)

func initDB() {
	InitDB("root", "123456", "127.0.0.1:3306", "video")
}

func TestVideo_Insert(t *testing.T) {
	initDB()
	defer mysql.Close()

	v := &Video{
		Title:  "test",
		Path:   "./test.mp4",
		UserID: 1,
	}
	if err := v.Insert(); err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestGetVideos(t *testing.T) {
	initDB()
	defer mysql.Close()

	videos, err := GetVideos(0)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range videos {
		fmt.Printf("%#v\n", *v)
	}
}
