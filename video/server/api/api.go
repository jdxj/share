package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	user "github.com/jdxj/share/user/proto"
	"github.com/jdxj/share/video/remote"
	"github.com/jdxj/share/video/server/api/comm"
	v1 "github.com/jdxj/share/video/server/api/v1"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// api
func RegisterAPI(apiGroup *gin.RouterGroup) {
	apiGroup.POST("sessions", AddSession)

	// api/v1
	v1Group := apiGroup.Group("v1")
	v1Group.Use(CheckLogin)
	v1.RegisterAPIV1(v1Group)

	// api/v2
	//v2Group := apiGroup.Group("v2")
	//v2Group.Use(CheckLogin)
	//{
	//	v2.RegisterAPIV2(v2Group)
	//}
}

func AddSession(c *gin.Context) {
	loginInfo := new(comm.User)
	err := c.ShouldBind(loginInfo)
	if err != nil {
		logger.Errorf("ShouldBind: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	req := &user.RequestLogin{
		Name:     loginInfo.Name,
		Password: loginInfo.Password,
	}
	loginResp, err := remote.UserService.Login(context.TODO(), req)
	if err != nil {
		logger.Errorf("Login: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if loginResp.Code != 0 {
		resp := comm.NewResponse(int(loginResp.Code), loginResp.Message, nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	uc := comm.NewUserClaims(int(loginResp.UserId), loginInfo.Name)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	secret, _ := KeyFunc(nil)
	ss, err := token.SignedString(secret)
	if err != nil {
		logger.Errorf("SignedString: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := comm.NewResponse(0, "token ok", ss)
	c.JSON(http.StatusOK, resp)
}

func CheckLogin(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		resp := comm.NewResponse(123, "not login", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	tokenStr := ExtractToken(bearerToken)
	uc := &comm.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, uc, KeyFunc)
	if err != nil {
		logger.Debugf("token: %s", tokenStr)
		logger.Errorf("ParseWithClaims: %s", err)
		resp := comm.NewResponse(123, "invalid token", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	uc = token.Claims.(*comm.UserClaims)
	c.Set(comm.UserState, uc)
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
