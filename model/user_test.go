package model

import "testing"

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
