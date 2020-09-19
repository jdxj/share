package v1

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	"github.com/jdxj/share/video/server/api"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(r *gin.RouterGroup) {
	// .
	//r.POST("sessions", AddSession)

	// v1
	v1Group := r.Group("v1")
	{
		v1Group.GET("", Home)
		v1Group.Static("assets", config.Server().AssetsPath)
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

func CheckLogin(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		resp := api.NewResponse(123, "not login", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	tokenStr := ExtractToken(bearerToken)
	uc := &api.UserClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, uc, KeyFunc)
	if err != nil {
		logger.Debugf("token: %s", tokenStr)
		logger.Errorf("ParseWithClaims: %s", err)
		resp := api.NewResponse(123, "invalid token", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
	}
}

func ExtractToken(tok string) string {
	// 注意 BEARER 后有一个空格
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:]
	}
	return tok
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(config.Server().Secret), nil
}
