package main

import (
	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	email "github.com/jdxj/share/email/proto"
	"github.com/jdxj/share/model"
	user "github.com/jdxj/share/user/proto"
	"github.com/jdxj/share/video/remote"
	"github.com/jdxj/share/video/server"
	"github.com/micro/micro/v3/service"
)

func initBase() error {
	err := config.Init("/home/jdxj/workspace/share/config.yaml")
	if err != nil {
		return err
	}
	logger.NewPathMode(config.Log().Path, config.Mode())

	dbCfg := config.DB()
	return model.InitDB(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.DBName)
}

func stopBase() error {
	_ = model.CloseDB()
	logger.Sync()
	return nil
}

func main() {
	srv := service.New(
		service.Name("video"),
		service.Version("latest"),
	)
	srv.Init(
		service.BeforeStart(initBase),
		service.BeforeStart(server.StartServer),
		service.BeforeStop(server.StopServer),
		service.BeforeStop(stopBase),
	)

	remote.UserService = user.NewUserService("user", srv.Client())
	remote.EmailService = email.NewEmailService("email", srv.Client())

	if err := srv.Run(); err != nil {
		logger.Errorf("Run: %s", err)
	}
}
