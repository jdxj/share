package handler

import (
	"context"
	"testing"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	email "github.com/jdxj/share/email/proto"
)

func TestEmail_Send(t *testing.T) {
	config.Init("/home/jdxj/workspace/share/config.yaml")
	logger.NewPathMode(config.Log().Path, config.Mode())
	e := new(Email)

	req := &email.RequestEmail{
		Token:      "123",
		Subject:    "test",
		Recipients: []string{"985759262@qq.com"},
		Type:       1,
		Content:    []byte("hello world"),
	}
	e.Send(context.TODO(), req, &email.ResponseEmail{})
}
