package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/asim/go-micro/v3/errors"

	pb_email "github.com/jdxj/share/email/proto"
)

const (
	host = "smtp.qq.com"
	addr = host + ":587"
)

type Email struct{}

func (e *Email) SendRaw(ctx context.Context, req *pb_email.RequestEmail, resp *pb_email.ResponseEmail) error {
	// todo: 验证 token
	//if req.Token == "" {
	//	resp.Code = 123
	//	resp.Message = "invalid token"
	//	return nil
	//}
	//
	//emailCfg := config.Email()
	//sender := email.NewEmail()
	//
	//sender.From = fmt.Sprintf("share <%s>", emailCfg.User)
	//sender.To = req.Recipients
	//
	//sender.Subject = req.Subject
	//switch req.Type {
	//case 2:
	//	sender.HTML = req.Content
	//default:
	//	sender.Text = req.Content
	//}
	//
	//err := sender.Send(addr, smtp.PlainAuth("", emailCfg.User, emailCfg.Token, host))
	//if err != nil {
	//	logger.Errorf("sender.Send: %s", err)
	//	return nil
	//}
	//
	//resp.Code = 0
	//resp.Message = "send email ok"
	//return nil

	resp.Message = "abc"
	fmt.Printf("alalallllll\n")
	// 测试
	//return fmt.Errorf("test return err: %s\n", "haha")
	//return errors.InternalServerError("email", "test return err: %s\n", "haha")
	return errors.Unauthorized("abc", "ff%s", "dd")
	//return errors.BadRequest("abc", "ff%s", "dd")
}

func (e *Email) Send(ctx context.Context, req *pb_email.RequestEmail, stream pb_email.Email_SendStream) error {
	fmt.Printf("alalallllll\n")
	err := stream.SendMsg(&pb_email.ResponseEmail{
		Code:    123,
		Message: "abc",
		Data:    []byte("hello world"),
	})
	if err != nil {
		log.Printf("SendMsg: %s\n", err)
	}
	return errors.Unauthorized("abc", "ff%s", "dd")
}
