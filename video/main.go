package main

import (
	"github.com/jdxj/logger"
	"github.com/micro/micro/v3/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("video"),
		service.Version("latest"),
	)

	// Register handler
	//pb.RegisterVideoHandler(srv.Server(), new(handler.Video))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Errorf("Run: %s", err)
	}
}
