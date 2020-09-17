package v1

import (
	"net/http"
	"path/filepath"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/video/server/api"

	"github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Errorf("FormFile: %s", err)
		resp := &api.Response{
			Code:    123,
			Message: "invalid param",
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// todo: 对 file name 进行安全处理
	err = c.SaveUploadedFile(
		file,
		filepath.Join(".", file.Filename),
	)
	if err != nil {
		logger.Errorf("SaveUploadedFile: %s", err)
		resp := &api.Response{
			Code:    123,
			Message: "can not save file",
			Data:    nil,
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := &api.Response{
		Code:    0,
		Message: "upload ok",
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}

func ListVideo(c *gin.Context) {

}
