package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewUserClaims(id int, name string) UserClaims {
	return UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    "jdxj",
		},
		ID:   id,
		Name: name,
	}
}

type UserClaims struct {
	jwt.StandardClaims

	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Name     string `form:"name"`
	Password string `form:"pass"`
}
