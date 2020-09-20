package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
)

var (
	server *http.Server
)

func StartServer() error {
	serverCfg := config.Server()
	server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "", serverCfg.Port),
		Handler: newRouter(),
	}
	logger.Infof("initServer success")

	go func() {
		logger.Infof("ListenAndServe before")

		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Errorf("ListenAndServe: %s", err)
			return
		}

		logger.Infof("ListenAndServe after")
	}()

	return nil
}

func StopServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
