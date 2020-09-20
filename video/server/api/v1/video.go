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
	uuid "github.com/satori/go.uuid"
)

// curl -X POST http://localhost:49152/api/v1/videos -F "file=@./mv.mp4" -H "Content-Type: multipart/form-data"
func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Errorf("FormFile: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	prefix, err := filepath.Abs(config.Server().AssetsPath)
	if err != nil {
		logger.Errorf("Abs: %s", err)
		resp := api.NewResponse(123, "unknown err", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	fileName := uuid.NewV4().String()
	filePath := filepath.Join(prefix, fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		logger.Errorf("SaveUploadedFile: %s", err)
		resp := api.NewResponse(123, "can not save file", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	v := &model.Video{
		Title:  file.Filename,
		Path:   fileName,
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
	if pageStr == "" {
		resp := api.NewResponse(123, "miss 'page' query param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logger.Errorf("ListVide Atoi: %s", err)
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

	p := &api.Page{
		Num:     page,
		HasNext: len(videos) >= model.CountPage,
		Data:    videos,
	}
	resp := api.NewResponse(0, "video list", p)
	c.JSON(http.StatusOK, resp)
}

func GetVideo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("GetVideo Atoi: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	v, err := model.GetVideoByID(id)
	if err != nil {
		logger.Errorf("model.GetVideoByID: %s", err)
		resp := api.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := api.NewResponse(0, "video", v)
	c.JSON(http.StatusOK, resp)
}
