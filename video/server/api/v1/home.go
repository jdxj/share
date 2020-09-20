package v1

import (
	"net/http"

	"github.com/jdxj/share/video/server/api/comm"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	resp := &comm.Response{
		Code:    0,
		Message: "ok",
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}
