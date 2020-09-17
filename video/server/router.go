package server

import (
	"github.com/jdxj/share/config"
	v1 "github.com/jdxj/share/video/server/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.Mode())
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	apiGroup := r.Group("api")
	v1.RegisterAPI(apiGroup)

	return r
}
