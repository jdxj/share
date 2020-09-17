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

func initServer() {
	serverCfg := config.Server()
	server = &http.Server{
		Addr: fmt.Sprintf("%s:%s", "127.0.0.1", serverCfg.Port),
		//Handler: NewRouter(),
	}
}

func StartServer() (err error) {
	initServer()
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
	return
}

func StopServer() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
