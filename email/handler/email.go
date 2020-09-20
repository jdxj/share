package handler

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/jdxj/logger"

	"github.com/jdxj/share/config"

	pb_email "github.com/jdxj/share/email/proto"

	"github.com/jordan-wright/email"
)

const (
	host = "smtp.qq.com"
	addr = host + ":587"
)

type Email struct{}

func (e *Email) Send(ctx context.Context, req *pb_email.RequestEmail, resp *pb_email.ResponseEmail) error {
	// todo: 验证 token
	if req.Token == "" {
		resp.Code = 123
		resp.Message = "invalid token"
		return nil
	}

	emailCfg := config.Email()
	sender := email.NewEmail()

	sender.From = fmt.Sprintf("share <%s>", emailCfg.User)
	sender.To = req.Recipients

	sender.Subject = req.Subject
	switch req.Type {
	case 2:
		sender.HTML = req.Content
	default:
		sender.Text = req.Content
	}

	err := sender.Send(addr, smtp.PlainAuth("", emailCfg.User, emailCfg.Token, host))
	if err != nil {
		logger.Errorf("sender.Send: %s", err)
		return nil
	}

	resp.Code = 0
	resp.Message = "send email ok"
	return nil
}
