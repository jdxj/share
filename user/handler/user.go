package handler

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/model"
	pb "github.com/jdxj/share/user/proto"
)

type User struct{}

func (u *User) Login(ctx context.Context, req *pb.RequestLogin, resp *pb.ResponseLogin) error {
	userInfo, err := model.LoginCheck(req.Name, DigPassword(req.Password))
	if err == nil {
		resp.Message = "ok"
		resp.UserId = int64(userInfo.ID)
		return nil
	}

	if err != sql.ErrNoRows {
		logger.Errorf("LoginCheck: %s", err)
		resp.Code = 101
		resp.Message = "internal error"
		return nil
	}

	resp.Code = 102
	resp.Message = "name or password error"
	return nil
}

func (u *User) SignUp(ctx context.Context, req *pb.RequestLogin, resp *pb.ResponseLogin) error {
	if req.Password == "" {
		resp.Code = 123
		resp.Message = "password can not empty"
		return nil
	}

	// todo: 是否加盐?
	userInfo := &model.User{
		Name:     req.Name,
		Password: DigPassword(req.Password),
	}
	err := userInfo.Insert()
	if err == nil {
		resp.Code = 0
		resp.Message = "new user ok"
		resp.UserId = userInfo.ID
		return nil
	}

	if err == model.ErrDuplicateName {
		resp.Code = 123
		resp.Message = model.ErrDuplicateName.Error()
		return nil
	}

	logger.Errorf("userInfo.Insert: %s", err)
	resp.Code = 123
	resp.Message = "internal err"
	return nil
}

func DigPassword(raw string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(raw)))
}
