package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jdxj/share/config"
)

// v1
func RegisterAPIV1(v1Group *gin.RouterGroup) {
	v1Group.GET("", Home)
	v1Group.Static("assets", config.Server().AssetsPath)

	// v1/videos
	videosGroup := v1Group.Group("videos")
	// videosGroup.Use()
	{
		videosGroup.POST("", UploadVideo)
		videosGroup.GET("", ListVideo)
		videosGroup.GET(":id", GetVideo)
	}
}
