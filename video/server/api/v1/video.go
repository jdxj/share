package v1

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	"github.com/jdxj/share/model"
	"github.com/jdxj/share/video/server/api"

	"github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Errorf("FormFile: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// todo: 对 file name 进行安全处理
	prefix, err := filepath.Abs(config.Server().AssetsPath)
	if err != nil {
		logger.Errorf("Abs: %s", err)
		resp := api.NewResponse(123, "unknow err", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	filePath := filepath.Join(prefix, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		logger.Errorf("SaveUploadedFile: %s", err)
		resp := api.NewResponse(123, "can not save file", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	v := &model.Video{
		Title:  file.Filename,
		Path:   filePath,
		UserID: 1, // todo: 权鉴后使用上传者 id
	}
	err = v.Insert()
	if err != nil {
		logger.Errorf("v.Insert: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := api.NewResponse(0, "upload ok", nil)
	c.JSON(http.StatusOK, resp)
}

func ListVideo(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logger.Errorf("Atoi: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	videos, err := model.GetVideos(page)
	if err != nil {
		logger.Errorf("GetVideos: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := api.NewResponse(0, "video list", videos)
	c.JSON(http.StatusOK, resp)
}

func GetVideo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("Atoi: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	v, err := model.GetVideo(id)
	if err != nil {
		logger.Errorf("model.GetVideo: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	fileName := filepath.Base(v.Path)
	resp := api.NewResponse(0, "video", fileName)
	c.JSON(http.StatusOK, resp)
}
