package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jdxj/share/config"
	"github.com/jdxj/share/video/server/api"
)

func newRouter() *gin.Engine {
	gin.SetMode(config.Mode())
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	apiGroup := r.Group("api")
	api.RegisterAPI(apiGroup)
	return r
}
