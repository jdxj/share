package v1

import (
	"context"
	"net/http"

	"github.com/jdxj/logger"
	user "github.com/jdxj/share/user/proto"
	"github.com/jdxj/share/video/server/api"
	"github.com/jdxj/share/video/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AddSession(c *gin.Context) {
	loginInfo := new(api.User)
	err := c.ShouldBind(loginInfo)
	if err != nil {
		logger.Errorf("ShouldBind: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	req := &user.RequestLogin{
		Name:     loginInfo.Name,
		Password: loginInfo.Password,
	}
	loginResp, err := service.UserService.Login(context.TODO(), req)
	if err != nil {
		logger.Errorf("Login: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if loginResp.Code != 0 {
		resp := api.NewResponse(int(loginResp.Code), loginResp.Message, nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	uc := api.NewUserClaims(int(loginResp.UserId), loginInfo.Name)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	secret, _ := KeyFunc(nil)
	ss, err := token.SignedString(secret)
	if err != nil {
		logger.Errorf("SignedString: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := api.NewResponse(0, "token ok", ss)
	c.JSON(http.StatusOK, resp)
}
