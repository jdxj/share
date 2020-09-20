package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jdxj/share/config"
)

// v1
func RegisterAPIV1(v1Group *gin.RouterGroup) {
	v1Group.GET("", Home)

	assetsGroup := v1Group.Group("assets")
	assetsGroup.Use(VerifyPermissions)
	{
		assetsGroup.Static("", config.Server().AssetsPath)
	}

	// v1/videos
	videosGroup := v1Group.Group("videos")
	// videosGroup.Use()
	{
		videosGroup.POST("", UploadVideo)
		videosGroup.GET("", ListVideo)
		videosGroup.GET(":id", GetVideo)
	}
}

// todo: 验证权限
func VerifyPermissions(c *gin.Context) {
}
