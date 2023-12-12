package controller

import (
	"Img/config"
	"Img/databases"
	"Img/model"
	"Img/services"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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
	fileName := file.Filename
	md5String := hex.EncodeToString(hash.Sum(nil))
	media := model.NewMedia()
	media.Md5 = md5String
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

	media.Name = fileName
	path := config.AppConf.StorageConf.Path + fileName
	media.Path = path
	err = c.SaveUploadedFile(file, path)
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

		fileName := file.Filename
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
