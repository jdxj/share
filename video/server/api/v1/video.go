package v1

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	email "github.com/jdxj/share/email/proto"
	"github.com/jdxj/share/model"
	"github.com/jdxj/share/video/remote"
	"github.com/jdxj/share/video/server/api/comm"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// curl -X POST http://localhost:49152/api/v1/videos -F "file=@./mv.mp4" -H "Content-Type: multipart/form-data"
func UploadVideo(c *gin.Context) {
	userState, _ := c.Get(comm.UserState)
	uc := userState.(*comm.UserClaims)
	logger.Debugf("id: %d, name: %s", uc.ID, uc.Name)

	file, err := c.FormFile("file")
	if err != nil {
		logger.Errorf("FormFile: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	prefix, err := filepath.Abs(config.Server().AssetsPath)
	if err != nil {
		logger.Errorf("Abs: %s", err)
		resp := comm.NewResponse(123, "unknown err", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	fileName := uuid.NewV4().String()
	filePath := filepath.Join(prefix, fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		logger.Errorf("SaveUploadedFile: %s", err)
		resp := comm.NewResponse(123, "can not save file", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	v := &model.Video{
		Title:  file.Filename,
		Path:   fileName,
		UserID: uc.ID,
	}
	err = v.Insert()
	if err != nil {
		logger.Errorf("v.Insert: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	reqEmail := &email.RequestEmail{
		Token:      "123",
		Subject:    "new video",
		Recipients: []string{"985759262@qq.com"}, // todo: 订阅者
		Type:       1,
		Content: []byte(fmt.Sprintf("%s:%s/api/v1/assets/%s",
			config.Server().Domain, config.Server().Port, fileName)),
	}

	// todo: rabbitmq?
	respEmail, err := remote.EmailService.Send(context.TODO(), reqEmail)
	if err != nil {
		logger.Errorf("EmailService.Send: %s", err)
		resp := comm.NewResponse(123, "send email failed", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := comm.NewResponse(0, "upload ok", respEmail)
	c.JSON(http.StatusOK, resp)
}

func ListVideo(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logger.Errorf("ListVide Atoi: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	videos, err := model.GetVideos(page)
	if err != nil {
		logger.Errorf("GetVideos: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	p := &comm.Page{
		Page:    page,
		HasNext: len(videos) >= model.CountPage,
		Data:    videos,
	}
	resp := comm.NewResponse(0, "video list", p)
	c.JSON(http.StatusOK, resp)
}

func GetVideo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Errorf("GetVideo Atoi: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	v, err := model.GetVideoByID(id)
	if err != nil {
		logger.Errorf("model.GetVideoByID: %s", err)
		resp := comm.NewResponse(123, "invalid param", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := comm.NewResponse(0, "video", v)
	c.JSON(http.StatusOK, resp)
}
