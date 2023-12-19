package controller

import (
	"Img/config"
	"Img/databases"
	"Img/model"
	"Img/services"
	"Img/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
	"time"
)

func Upload(c *gin.Context) {
	resp := services.NewSuccessResponse()
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	open, err := file.Open()
	if err != nil {
		return
	}
	hash := md5.New()
	_, _ = io.Copy(hash, open)
	err = open.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fileNameExt := filepath.Ext(file.Filename)
	md5String := hex.EncodeToString(hash.Sum(nil))
	//filePathName := time.Now().Format("2006-01-02") + util.GetPathTag() + md5String + fileNameExt
	filePathName := util.GetPathTag() + md5String + fileNameExt
	media := model.NewMedia()
	media.Md5 = md5String
	media.Size = file.Size
	media.Type = file.Header.Get("Content-Type")
	err = databases.GetMediaByMd5(media)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if media.ID > 0 {
		resp.Data = media
		c.JSON(http.StatusOK, resp)
		return
	}
	media.Name = file.Filename
	storagePath := config.AppConf.StorageConf.Path + filePathName
	media.Path = filePathName
	err = c.SaveUploadedFile(file, storagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = databases.CreateMedia(media)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, media)
}

func UploadMultiple(c *gin.Context) {
	resp := services.NewSuccessResponse()
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	files := form.File["files"]
	var mediaModels []*model.Media
	for _, file := range files {
		open, err := file.Open()
		if err != nil {
			return
		}
		hash := md5.New()
		_, _ = io.Copy(hash, open)
		err = open.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		fileName := time.Now().Format("2006-01-02") + util.GetPathTag() + file.Filename
		path := config.AppConf.StorageConf.Path + fileName
		md5String := hex.EncodeToString(hash.Sum(nil))
		media := model.NewMedia()
		media.Md5 = md5String
		media.Path = path
		media.Name = fileName
		err = databases.GetMediaByMd5(media)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if media.ID > 0 {
			resp.Data = media
			mediaModels = append(mediaModels, media)
			continue
		}
		media.Path = path
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		err = databases.CreateMedia(media)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		mediaModels = append(mediaModels, media)
	}
	resp.Data = mediaModels
	c.JSON(http.StatusOK, resp)
}
