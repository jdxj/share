package handler

import (
	"fmt"
	"testing"

	"github.com/jdxj/share/model"

	user "github.com/jdxj/share/user/proto"
)

func TestUser_Login(t *testing.T) {
	var a, b uint32 = 3, 7
	fmt.Printf("%d\n", int(a-b))
	var c, d int32 = 3, 7
	fmt.Printf("%d\n", c-d)
}

func TestUser_SignUp(t *testing.T) {
	model.InitDB("root", "123456", "127.0.0.1:3306", "video")
	defer model.CloseDB()

	u := new(User)
	req := &user.RequestLogin{
		Name:     "abc",
		Password: "jdxj",
	}
	resp := &user.ResponseLogin{}
	_ = u.SignUp(nil, req, resp)
	fmt.Printf("%#v\n", resp)
}
