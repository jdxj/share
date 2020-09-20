package main

import (
	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	"github.com/jdxj/share/email/handler"
	pb "github.com/jdxj/share/email/proto"
	"github.com/micro/micro/v3/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("email"),
		service.Version("latest"),
	)
	srv.Init(
		service.BeforeStart(initBase),
	)

	// Register handler
	pb.RegisterEmailHandler(srv.Server(), new(handler.Email))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Errorf("srv.Run: %s", err)
	}
}

func initBase() error {
	err := config.Init("/home/jdxj/workspace/share/config.yaml")
	if err != nil {
		return err
	}
	logger.NewPathMode(config.Log().Path, config.Mode())
	return nil
}
